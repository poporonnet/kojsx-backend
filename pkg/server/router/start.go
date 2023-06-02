package router

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/handlers"
	"github.com/mct-joken/kojs5-backend/pkg/utils/mail/dummy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	contestHandler *handlers.ContestHandlers
	userHandler    *handlers.UserHandlers
	problemHandler *handlers.ProblemHandlers
	logger         *zap.Logger
	mongoClient    *mongodb.Client
)

func initServer() {
	mode := os.Getenv("KOJS_MODE")
	var (
		contestRepository repository.ContestRepository
		userRepository    repository.UserRepository
		problemRepository repository.ProblemRepository
	)

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ = cfg.Build()

	if mode == "prod" {
		mongoClient = mongodb.NewMongoDBClient("mongodb://localhost:27017")
		contestRepository = mongodb.NewContestRepository(*mongoClient)
		userRepository = mongodb.NewUserRepository(*mongoClient)
		problemRepository = mongodb.NewProblemRepository(*mongoClient)
		logger.Sugar().Info("start the server in production mode.")
	} else {
		contestRepository = inmemory.NewContestRepository([]domain.Contest{})
		userRepository = inmemory.NewUserRepository([]domain.User{})
		problemRepository = inmemory.NewProblemRepository([]domain.Problem{}, []domain.Caseset{}, []domain.Case{})
		logger.Sugar().Info("start the server in development mode.")
	}

	contestHandler = handlers.NewContestHandlers(
		*controller.NewContestController(
			contestRepository,
			*contest.NewCreateContestService(contestRepository),
			*contest.NewFindContestService(contestRepository),
		),
		logger,
	)

	userHandler = handlers.NewUserHandlers(
		*controller.NewUserController(
			userRepository,
			*user.NewCreateUserService(
				userRepository,
				*service.NewUserService(userRepository),
				dummy.NewMailer(),
				"",
			),
			*user.NewFindUserService(userRepository),
		),
		*controller.NewAuthController(userRepository, ""),
		logger,
	)

	problemHandler = handlers.NewProblemHandlers(
		*controller.NewProblemController(
			problemRepository,
			*problem.NewCreateProblemService(
				problemRepository,
				*service.NewProblemService(problemRepository),
			),
			*problem.NewFindProblemService(problemRepository),
		),
		logger,
	)
}

func StartServer(port int) {
	initServer()
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
	if mongoClient != nil {
		if err := mongoClient.Cli.Disconnect(ctx); err != nil {
			logger.Sugar().Fatal(err)
		}
		logger.Sugar().Info("Disconnected from database.")
	}

}
