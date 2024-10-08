package server

import (
	"log"

	"github.com/phetployst/sekai-shop-microservices/modules/item/itemHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemPb"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler

	item := s.app.Group("/item_v1")

	item.GET("/", s.healthCheckService)

}
