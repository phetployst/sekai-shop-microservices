package itemHandler

import "github.com/phetployst/sekai-shop-microservices/modules/item/itemUsecase"

type itemGrpcHandler struct {
	itemUsecase itemUsecase.ItemUsecaseService
}

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
