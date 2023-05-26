package store

import (
	product2 "PowerX/internal/model/product"
	"PowerX/internal/uc/powerx/product"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListStoresPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStoresLogic {
	return &ListStoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListStoresLogic) ListStoresPage(req *types.ListStoresPageRequest) (resp *types.ListStoresPageReply, err error) {
	stores, err := l.svcCtx.PowerX.Store.FindManyStores(l.ctx, &product.FindManyStoresOption{
		LikeName: req.LikeName,
		OrderBy:  req.OrderBy,
		PageEmbedOption: types.PageEmbedOption{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
		},
	})

	if err != nil {
		return nil, err
	}
	list := TransferStoresToStoresReply(stores.List)

	return &types.ListStoresPageReply{
		List:      list,
		PageIndex: stores.PageIndex,
		PageSize:  stores.PageSize,
		Total:     stores.Total,
	}, nil

}

func TransferStoresToStoresReply(stores []*product2.Store) []*types.Store {
	storesReply := []*types.Store{}
	for _, store := range stores {
		storeReply := TransferStoreToStoreReply(store)
		storesReply = append(storesReply, storeReply)
	}
	return storesReply
}

func TransferStoreToStoreReply(artisan *product2.Store) *types.Store {
	return &types.Store{
		Id:            artisan.Id,
		Name:          artisan.Name,
		EmployeeId:    artisan.EmployeeId,
		ContactNumber: artisan.ContactNumber,
		CoverURL:      artisan.CoverURL,
		Address:       artisan.Address,
		Longitude:     artisan.Longitude,
		Latitude:      artisan.Latitude,
		CreatedAt:     artisan.CreatedAt.String(),
		Artisans:      TransferArtisansToShopArtisans(artisan.Artisans),
	}
}
