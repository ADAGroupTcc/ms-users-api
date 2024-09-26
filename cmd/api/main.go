package main

import (
	"context"

	"github.com/ADAGroupTcc/ms-users-api/config"
	"github.com/ADAGroupTcc/ms-users-api/internal/http/router"
	"github.com/ADAGroupTcc/ms-users-api/pkg/start"
)

func main() {
	ctx := context.TODO()
	envs, err := config.LoadEnvVars()
	if err != nil {
		panic(err)
	}
	dependencies := config.NewDependencies(ctx, envs)
	e := router.SetupRouter(dependencies)
	if err := start.StartServer(e, envs.ApiPort, dependencies.Database); err != nil {
		panic(err)
	}
}
