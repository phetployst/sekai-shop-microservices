package paymentHandler

import (
	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
