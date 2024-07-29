package config

import (
	"context"

	"github.com/ADAGroupTcc/ms-users-api/internal/http/health"
	handler "github.com/ADAGroupTcc/ms-users-api/internal/http/users"
	repository "github.com/ADAGroupTcc/ms-users-api/internal/repositories/users"
	healthService "github.com/ADAGroupTcc/ms-users-api/internal/services/health"
	service "github.com/ADAGroupTcc/ms-users-api/internal/services/users"
	"github.com/ADAGroupTcc/ms-users-api/pkg/mongorm"
)

type Dependencies struct {
	Handler       handler.Handler
	HealthHandler health.Health
}

func NewDependencies(ctx context.Context, envs *Environments) *Dependencies {
	database, err := mongorm.Connect(envs.DBUri, envs.DBName)
	if err != nil {
		panic(err)
	}
	userRepository := repository.New(database)
	userService := service.New(userRepository)
	userHandler := handler.New(userService)

	healthService := healthService.New(database)
	healthHandler := health.New(healthService)
	return &Dependencies{
		userHandler,
		healthHandler,
	}
}
