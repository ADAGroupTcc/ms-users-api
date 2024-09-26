package start

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ADAGroupTcc/ms-users-api/pkg/mongorm"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartServer(e *echo.Echo, apiPort string, connection *mongo.Database) error {
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", apiPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Erro ao iniciar o servidor:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Println("Stopping server gracefully...")
	defer cancel()
	err := mongorm.Close(connection)
	if err != nil {
		fmt.Println("Erro ao fechar a conexão com o banco de dados:", err)
	}

	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Erro ao desligar o servidor:", err)
	}
	return nil
}
