package powerx

import (
	"PowerX/internal/config"
	"PowerX/internal/model/media"
	"PowerX/pkg/filex"
	"PowerX/pkg/httpx"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"path/filepath"
)

const DefaultStoragePath = "resource/static"

// MediaResourceUseCase Use Case
type MediaResourceUseCase struct {
	db               *gorm.DB
	OSSClient        *minio.Client
	LocalStoragePath string
	LocalStorageUrl  string
}

const BucketProduct = "bucket.product"

func NewMediaResourceUseCase(db *gorm.DB, conf *config.Config) *MediaResourceUseCase {
	// 使用Minio API SDK
	var c *minio.Client
	var err error
	localStoragePath := DefaultStoragePath
	localStorageUrl, _ := httpx.GetURL(conf.Server.Host, conf.Server.Port, conf.MediaResource.LocalStorage.StoragePath)
	if conf.MediaResource.OSS.Enable {
		minioConfig := conf.MediaResource.OSS.Minio
		c, err = minio.New(minioConfig.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(minioConfig.Credentials.AccessKey, minioConfig.Credentials.SecretKey, ""),
			Secure: minioConfig.UseSSL,
		})

		if err != nil {
			panic(errors.Wrap(err, "minio client init failed"))
		}
	} else {
		if conf.MediaResource.LocalStorage.StoragePath != "" {
			localStoragePath = filepath.Join(localStoragePath, conf.MediaResource.LocalStorage.StoragePath)
		}
	}

	return &MediaResourceUseCase{
		db:               db,
		OSSClient:        c,
		LocalStoragePath: localStoragePath,
		LocalStorageUrl:  localStorageUrl,
	}
}

func (uc *MediaResourceUseCase) CreateMediaResource(ctx context.Context, store *media.MediaResource) {
	if err := uc.db.WithContext(ctx).Create(&store).Error; err != nil {
		panic(err)
	}
}

func (uc *MediaResourceUseCase) MakeProductMediaResource(ctx context.Context, handle *multipart.FileHeader) (resource *media.MediaResource, err error) {
	return uc.MakeMediaResource(ctx, BucketProduct, handle)
}
func (uc *MediaResourceUseCase) MakeMediaResource(ctx context.Context, bucket string, handle *multipart.FileHeader) (resource *media.MediaResource, err error) {

	if uc.OSSClient != nil {
		resource, err = uc.MakeOSSResource(ctx, bucket, handle)
	} else {
		resource, err = uc.MakeLocalResource(ctx, bucket, handle)
	}

	if err != nil {
		return nil, err
	}

	uc.CreateMediaResource(ctx, resource)

	return
}

func (uc *MediaResourceUseCase) MakeLocalResource(ctx context.Context, bucket string, handle *multipart.FileHeader) (resource *media.MediaResource, err error) {

	// 获取文件名和文件大小
	filename := handle.Filename
	filesize := handle.Size

	// 模拟将文件保存到本地存储的逻辑
	// 这里可以根据实际需求进行处理，比如将文件保存到指定的文件夹或存储路径中
	// 以下示例将文件保存到名为 "uploads" 的文件夹下
	bucketPath := filepath.Join(uc.LocalStoragePath, bucket)
	// 检查目录是否存在
	if _, err = os.Stat(bucketPath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err = os.MkdirAll(bucketPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	uploadPath := filepath.Join(bucketPath, filename)
	err = filex.SaveFileToLocal(handle, uploadPath)
	if err != nil {
		return nil, err
	}

	contentType := handle.Header.Get("Content-Type")

	url, err := httpx.AppendURIs(uc.LocalStorageUrl, uploadPath)
	if err != nil {
		return nil, err
	}

	// 构建媒体资源对象
	resource = &media.MediaResource{
		Filename:     filename,
		Size:         filesize,
		Url:          url,
		ContentType:  handle.Header.Get("Content-Type"),
		ResourceType: filex.GetFileType(contentType),
	}

	return resource, nil
}

func (uc *MediaResourceUseCase) MakeOSSResource(ctx context.Context, bucket string, handle *multipart.FileHeader) (resource *media.MediaResource, err error) {
	exist, err := uc.OSSClient.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}

	if !exist {
		location := "us-east-1"
		err = uc.OSSClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return nil, err
		}
		// 设置存储桶策略
		policy := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": "*",
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::` + bucket + `/*"]
			}
		]
	}`
		err = uc.OSSClient.SetBucketPolicy(context.Background(), bucket, policy)
		if err != nil {
			return nil, err
		}
	}

	// Upload the resource file
	objectName := handle.Filename
	filePath, _ := handle.Open()
	contentType := handle.Header.Get("Content-Type")
	info, err := uc.OSSClient.PutObject(ctx, bucket, objectName, filePath, handle.Size, minio.PutObjectOptions{ContentType: contentType})

	url, _ := httpx.AppendURIs(uc.OSSClient.EndpointURL().String(), bucket, info.Key)

	resource = &media.MediaResource{
		Filename:     info.Key,
		Size:         info.Size,
		Url:          url,
		ContentType:  contentType,
		ResourceType: filex.GetFileType(contentType),
	}

	return resource, err
}
