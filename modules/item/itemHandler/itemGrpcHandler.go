package itemHandler

import (
	"context"

	"github.com/phetployst/sekai-shop-microservices/modules/item/itemPb"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemUsecase"
)

type itemGrpcHandler struct {
	itemPb.UnimplementedItemGrpcServiceServer
	itemUsecase itemUsecase.ItemUsecaseService
}

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}

func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}
