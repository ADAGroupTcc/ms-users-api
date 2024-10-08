package router

import (
	"github.com/ADAGroupTcc/ms-users-api/config"
	"github.com/ADAGroupTcc/ms-users-api/internal/http/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRouter(dependencies *config.Dependencies) *echo.Echo {
	e := echo.New()

	e.GET("/health", dependencies.HealthHandler.Check)

	v1 := e.Group("/v1")
	v1.POST("/users", dependencies.Handler.Create, middlewares.ErrorIntercepter())
	v1.GET("/users/:id", dependencies.Handler.Get, middlewares.ErrorIntercepter())
	v1.GET("/users", dependencies.Handler.List, middlewares.ErrorIntercepter())
	v1.PATCH("/users/:id", dependencies.Handler.Update, middlewares.ErrorIntercepter())
	v1.DELETE("/users/:id", dependencies.Handler.Delete, middlewares.ErrorIntercepter())

	return e
}
