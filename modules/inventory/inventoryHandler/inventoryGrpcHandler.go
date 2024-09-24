package inventoryHandler

import (
	"github.com/phetployst/sekai-shop-microservices/modules/inventory/inventoryUsecase"
)

type inventoryGrpcHandler struct {
	inventoryUsecase inventoryUsecase.InventoryUsecaseService
}

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
