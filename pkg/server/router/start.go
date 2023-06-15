package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/utils"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/application/submission"
	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/handlers"
	"github.com/mct-joken/kojs5-backend/pkg/utils/mail/dummy"
	"go.uber.org/zap"
)

var (
	contestHandler    *handlers.ContestHandlers
	userHandler       *handlers.UserHandlers
	problemHandler    *handlers.ProblemHandlers
	submissionHandler *handlers.SubmissionHandlers
	logger            *zap.Logger
	mongoClient       *mongodb.Client
)

func initServer(mode string) {
	var (
		contestRepository    repository.ContestRepository
		userRepository       repository.UserRepository
		problemRepository    repository.ProblemRepository
		submissionRepository repository.SubmissionRepository
	)

	logger = utils.NewLogger()
	if mode == "prod" {
		utils.SugarLogger.Info("start the server in production mode.")
		utils.SugarLogger.Info("connect to mongodb server...")
		mongoClient = mongodb.NewMongoDBClient("mongodb://localhost:27017")
		contestRepository = mongodb.NewContestRepository(*mongoClient)
		userRepository = mongodb.NewUserRepository(*mongoClient)
		problemRepository = mongodb.NewProblemRepository(*mongoClient)
		submissionRepository = mongodb.NewSubmissionRepository(*mongoClient)
		utils.SugarLogger.Info("repository initialized")
	} else {
		logger.Sugar().Info("start the server in development mode.")
		utils.SugarLogger.Info("in-memory repository initialisation with seeds...")
		seeds := seed.NewSeeds()
		contestRepository = inmemory.NewContestRepository(seeds.Contests)
		userRepository = inmemory.NewUserRepository(seeds.Users)
		problemRepository = inmemory.NewProblemRepository(seeds.Problems)
		submissionRepository = inmemory.NewSubmissionRepository(seeds.Submission)
		utils.SugarLogger.Info("repository initialized")
	}

	contestHandler = func() *handlers.ContestHandlers {
		createService := *contest.NewCreateContestService(contestRepository)
		findService := *contest.NewFindContestService(contestRepository)
		c := *controller.NewContestController(
			contestRepository,
			createService,
			findService,
		)
		return handlers.NewContestHandlers(
			c,
			logger,
		)
	}()

	userHandler = func() *handlers.UserHandlers {
		findService := *user.NewFindUserService(userRepository)
		createService := *user.NewCreateUserService(
			userRepository,
			*service.NewUserService(userRepository),
			dummy.NewMailer(),
			"",
		)

		cont := *controller.NewUserController(
			userRepository,
			createService,
			findService,
		)

		auth := *controller.NewAuthController(userRepository, "")

		return handlers.NewUserHandlers(
			cont,
			auth,
			logger,
		)
	}()

	problemHandler = func() *handlers.ProblemHandlers {
		createService := *problem.NewCreateProblemService(
			problemRepository,
			*service.NewProblemService(problemRepository),
		)
		cont := *controller.NewProblemController(
			problemRepository,
			createService,
			*problem.NewFindProblemService(problemRepository, contestRepository),
		)

		return handlers.NewProblemHandlers(
			cont,
			logger,
		)
	}()

	submissionHandler = func() *handlers.SubmissionHandlers {
		createService := *submission.NewCreateSubmissionService(
			submissionRepository,
			*service.NewSubmissionService(submissionRepository),
			problemRepository,
		)
		findService := *submission.NewFindSubmissionService(submissionRepository, problemRepository)
		findProblemService := *problem.NewFindProblemService(problemRepository, contestRepository)
		findUserService := *user.NewFindUserService(userRepository)
		cont := *controller.NewSubmissionController(
			submissionRepository,
			createService,
			findService,
			findProblemService,
			findUserService,
		)
		return handlers.NewSubmissionHandlers(cont, logger)
	}()
}

func StartServer(port int, mode string) {
	initServer(mode)
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
	if mongoClient != nil {
		if err := mongoClient.Cli.Disconnect(ctx); err != nil {
			logger.Sugar().Fatal(err)
		}
		logger.Sugar().Info("Disconnected from database.")
	}
}
