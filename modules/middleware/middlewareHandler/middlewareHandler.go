package middlewareHandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/phetployst/sekai-shop-microservices/config"
	"github.com/phetployst/sekai-shop-microservices/modules/middleware/middlewareUsecase"
	"github.com/phetployst/sekai-shop-microservices/pkg/response"
)

type (
	MiddlewareHandlerService interface {
		JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

func (h *middlewareHandler) JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		newCtx, err := h.middlewareUsecase.JwtAuthorization(c, h.cfg, accessToken)
		if err != nil {
			return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
		}

		return next(newCtx)
	}
}
