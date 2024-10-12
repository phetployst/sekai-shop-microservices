package server

import (
	"github.com/phetployst/sekai-shop-microservices/modules/inventory/inventoryHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/inventory/inventoryRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	inventory.GET("/", s.healthCheckService)
	inventory.GET("/inventory/:player_id", httpHandler.FindPlayerItems, s.middleware.JwtAuthorization, s.middleware.PlayerIdParamValidation)
}
