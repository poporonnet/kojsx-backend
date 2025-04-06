package user

import (
	"github.com/poporonnet/kojsx-backend/pkg/server"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/handlers"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/domainService"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"
)

var (
	UserHandler *handlers.UserHandlers
)

func init() {
	// ToDo: Userモジュールで必要になるものの初期化をやる
	logger := utils.NewLogger()
	seeds := seed.NewSeeds()

	userRepository := inmemory.NewUserRepository(seeds.Users)

	findService := *service.NewFindUserService(userRepository)
	createService := *service.NewCreateUserService(userRepository,
		*domainService.NewUserService(userRepository),
		dummy.NewMailer(),
		"",
	)

	ctrl := *controller.NewUserController(
		userRepository,
		createService,
		findService,
	)

	UserHandler = handlers.NewUserHandlers(
		ctrl,
		*server.NewAuthController(userRepository, ""),
		logger,
	)
}
