package itemHandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/item"
	"github.com/phetployst/sekai-shop-microservices/modules/item/itemUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/request"
	"github.com/phetployst/sekai-shop-microservices/pkg/response"
)

type (
	ItemHttpHandlerService interface {
		CreateItem(c echo.Context) error
	}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUsecase: itemUsecase,
	}
}

func (h *itemHttpHandler) CreateItem(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(item.CreateItemReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.itemUsecase.CreateItem(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}
