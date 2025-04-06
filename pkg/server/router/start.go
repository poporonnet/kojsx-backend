package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/poporonnet/kojsx-backend/pkg/contest"
	"github.com/poporonnet/kojsx-backend/pkg/user"
	"go.uber.org/zap"
)

var (
	contestHandler    = contest.ContestHandler
	userHandler       = user.UserHandler
	problemHandler    = contest.ProblemHandler
	submissionHandler = contest.SubmissionHandler
	logger            *zap.Logger
)

func StartServer(port int, mode string) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 5 << 10,
		LogLevel:  1,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger.Sugar().Errorf("ERROR: %v\n%s\n", err, string(stack))
			return err
		},
	}))
	e.HideBanner = true
	e.HidePort = true
	// routerの呼び出し
	rootRouter(e)

	// Ctrl+Cやkill signalを受け取る
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	go func() {
		if err := e.Start(fmt.Sprintf("localhost:%d", port)); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, stopper := context.WithTimeout(context.Background(), 5*time.Second)
	defer stopper()

	logger.Sugar().Info("Server shutdown has begun.")
	if err := e.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal(err)
	}
	logger.Sugar().Info("WebServer terminated successfully.")
}
