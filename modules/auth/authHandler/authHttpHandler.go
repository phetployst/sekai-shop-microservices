package authHandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/auth"
	"github.com/phetployst/sekai-shop-microservices/modules/auth/authUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/request"
	"github.com/phetployst/sekai-shop-microservices/pkg/response"
)

type (
	AuthHttpHandlerService interface {
		Login(c echo.Context) error
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}

func (h *authHttpHandler) Login(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(auth.PlayerLoginReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Login(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
