package contest

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/handlers"
	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/service"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	userInmemory "github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	userService "github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"
)

var (
	ContestHandler    *handlers.ContestHandlers
	ProblemHandler    *handlers.ProblemHandlers
	SubmissionHandler *handlers.SubmissionHandlers
)

func init() {
	logger := utils.NewLogger()
	seeds := seed.NewSeeds()

	// Repositoryの初期化
	contestRepository := inmemory.NewContestRepository(seeds.Contests)
	contestantRepository := inmemory.NewContestantRepository(seeds.Contestants)
	problemRepository := inmemory.NewProblemRepository(seeds.Problems)
	submissionRepository := inmemory.NewSubmissionRepository(seeds.Submission)
	// ToDo: userRepositoryを直接操作するのをやめて，何らかのモジュール間通信を利用する
	userRepository := userInmemory.NewUserRepository(seeds.Users)

	// ContestHandler
	createContestService := *contest.NewCreateContestService(contestRepository, contestantRepository, *service.NewContestantService(contestantRepository))
	findContestService := *contest.NewFindContestService(contestRepository)
	rankingService := contest.NewGetContestRankingService(
		contestRepository,
		contestantRepository,
		problemRepository,
		submissionRepository,
		userRepository,
	)
	contestCtrl := *controller.NewContestController(
		contestRepository,
		createContestService,
		findContestService,
		*rankingService,
	)
	ContestHandler = handlers.NewContestHandlers(contestCtrl, logger)

	// ProblemHandler
	createProblemService := *problem.NewCreateProblemService(
		problemRepository,
		*service.NewProblemService(problemRepository),
	)
	problemCtrl := *controller.NewProblemController(
		problemRepository,
		createProblemService,
		*problem.NewFindProblemService(problemRepository, contestRepository, contestantRepository),
	)
	ProblemHandler = handlers.NewProblemHandlers(problemCtrl, logger)

	// SubmissionHandler
	createSubmission := *submission.NewCreateSubmissionService(
		submissionRepository,
		*service.NewSubmissionService(submissionRepository),
		problemRepository,
	)
	findSubmission := *submission.NewFindSubmissionService(submissionRepository, problemRepository)
	findProblem := *problem.NewFindProblemService(problemRepository, contestRepository, contestantRepository)
	findUserService := *userService.NewFindUserService(userRepository)
	cont := *controller.NewSubmissionController(
		submissionRepository,
		createSubmission,
		findSubmission,
		findProblem,
		findUserService,
	)
	SubmissionHandler = handlers.NewSubmissionHandlers(cont, logger)
}
