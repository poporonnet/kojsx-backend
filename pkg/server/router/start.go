package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	controller2 "github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller"
	handlers2 "github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/handlers"
	inmemory3 "github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	repository2 "github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	service2 "github.com/poporonnet/kojsx-backend/pkg/contest/model/service"
	contest2 "github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	problem2 "github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	submission2 "github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/handlers"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/service"
	service3 "github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"go.uber.org/zap"
)

var (
	contestHandler    *handlers2.ContestHandlers
	userHandler       *handlers.UserHandlers
	problemHandler    *handlers2.ProblemHandlers
	submissionHandler *handlers2.SubmissionHandlers
	logger            *zap.Logger
)

func initServer(mode string) {
	var (
		contestRepository    repository2.ContestRepository
		userRepository       repository.UserRepository
		problemRepository    repository2.ProblemRepository
		submissionRepository repository2.SubmissionRepository
		contestantRepository repository2.ContestantRepository
	)

	logger = utils.NewLogger()
	logger.Sugar().Info("start the server in development mode.")
	utils.SugarLogger.Info("in-memory repository initialisation with seeds...")
	seeds := seed.NewSeeds()
	contestRepository = inmemory3.NewContestRepository(seeds.Contests)
	userRepository = inmemory.NewUserRepository(seeds.Users)
	problemRepository = inmemory3.NewProblemRepository(seeds.Problems)
	submissionRepository = inmemory3.NewSubmissionRepository(seeds.Submission)
	utils.SugarLogger.Info("repository initialized")

	contestHandler = func() *handlers2.ContestHandlers {
		createService := *contest2.NewCreateContestService(contestRepository, contestantRepository, *service2.NewContestantService(contestantRepository))
		findService := *contest2.NewFindContestService(contestRepository)
		rankingService := contest2.NewGetContestRankingService(
			contestRepository,
			contestantRepository,
			problemRepository,
			submissionRepository,
			userRepository,
		)
		c := *controller2.NewContestController(
			contestRepository,
			createService,
			findService,
			*rankingService,
		)
		return handlers2.NewContestHandlers(
			c,
			logger,
		)
	}()

	userHandler = func() *handlers.UserHandlers {
		findService := *service3.NewFindUserService(userRepository)
		createService := *service3.NewCreateUserService(
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

		auth := *controller2.NewAuthController(userRepository, "")

		return handlers.NewUserHandlers(
			cont,
			auth,
			logger,
		)
	}()

	problemHandler = func() *handlers2.ProblemHandlers {
		createService := *problem2.NewCreateProblemService(
			problemRepository,
			*service2.NewProblemService(problemRepository),
		)
		cont := *controller2.NewProblemController(
			problemRepository,
			createService,
			*problem2.NewFindProblemService(problemRepository, contestRepository, contestantRepository),
		)

		return handlers2.NewProblemHandlers(
			cont,
			logger,
		)
	}()

	submissionHandler = func() *handlers2.SubmissionHandlers {
		createService := *submission2.NewCreateSubmissionService(
			submissionRepository,
			*service2.NewSubmissionService(submissionRepository),
			problemRepository,
		)
		findService := *submission2.NewFindSubmissionService(submissionRepository, problemRepository)
		findProblemService := *problem2.NewFindProblemService(problemRepository, contestRepository, contestantRepository)
		findUserService := *service3.NewFindUserService(userRepository)
		cont := *controller2.NewSubmissionController(
			submissionRepository,
			createService,
			findService,
			findProblemService,
			findUserService,
		)
		return handlers2.NewSubmissionHandlers(cont, logger)
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
}
