package product

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powermodel"
)

type ProductCategory struct {
	powermodel.PowerModel

	PId         int64  `gorm:"comment:上级品类"`
	Name        string `gorm:"comment:品类名称"`
	Sort        int8   `gorm:"comment:排序"`
	ViceName    string `gorm:"comment:副标题"`
	Description string `gorm:"comment:描述"`

	model.ImageAbleInfo
}

const ProductCategoryUniqueId = powermodel.UniqueId
