package server

import (
	"github.com/phetployst/sekai-shop-microservices/modules/payment/paymentHandler"
	"github.com/phetployst/sekai-shop-microservices/modules/payment/paymentRepository"
	"github.com/phetployst/sekai-shop-microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	_ = payment

}
