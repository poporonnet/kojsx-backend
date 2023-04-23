package router

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(port int) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.HideBanner = true

	// routerの呼び出し
	rootRouter(e)

	// グレイスフルシャットダウン用
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
			e.Logger.Fatal("Shutting down server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
