package router

import (
	"github.com/ADAGroupTcc/ms-users-api/config"
	"github.com/ADAGroupTcc/ms-users-api/internal/http/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRouter(dependencies *config.Dependencies) *echo.Echo {
	e := echo.New()
	e.Use(middlewares.ErrorIntercepter())

	e.GET("/health", dependencies.HealthHandler.Check)

	v1 := e.Group("/v1")
	v1.POST("/users", dependencies.Handler.Create)
	v1.GET("/users/:id", dependencies.Handler.Get)
	v1.GET("/users", dependencies.Handler.List)
	v1.PATCH("/users/:id", dependencies.Handler.Update)
	v1.DELETE("/users/:id", dependencies.Handler.Delete)

	return e
}
