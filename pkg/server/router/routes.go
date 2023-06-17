package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func rootRouter(e *echo.Echo) {
	e.GET("/", topHandler)
	v2 := e.Group("/api/v2")
	{
		v2.POST("/login", userHandler.Login)
		v2.POST("/logout", dummyHandler)
		v2.POST("/refresh", dummyHandler)

		user := v2.Group("/users")
		{
			user.GET("me", dummyHandler)
			user.PUT("me/password", dummyHandler)

			user.GET("/:id", userHandler.FindByID)
			user.GET("", userHandler.FindAllUser)
			user.POST("", userHandler.CreateUser)
			user.POST("verify/:token", userHandler.Verify)
		}

		problem := v2.Group("/problems")
		{
			problem.POST("/", problemHandler.CreateProblem)

			problem.GET("/:id", problemHandler.FindByID)
			problem.PUT("/:id", dummyHandler)

			problem.POST("/:id/sets", dummyHandler)
			problem.PUT("/:id/sets/:setId", dummyHandler)
			problem.DELETE("/:id/sets/:setId", dummyHandler)

			problem.POST("/:id/sets/:setId/cases", dummyHandler)
			problem.PUT("/:id/sets/:setId/cases/:caseId", dummyHandler)
			problem.DELETE("/:id/sets/:setId/cases/:caseId", dummyHandler)
		}

		contest := v2.Group("/contests")
		{
			contest.POST("/", contestHandler.CreateContest)
			contest.GET("/", contestHandler.FindContest)
			contest.GET("/:id", contestHandler.FindContestByID)
			contest.PUT("/:id", dummyHandler)
			contest.POST("/:id/join", dummyHandler)
			contest.GET("/:id/problems", problemHandler.FindByContestID)
			contest.GET("/:id/ranking", contestHandler.GetRanking)

			contest.POST("/:id/submissions", submissionHandler.CreateSubmission)
			contest.GET("/:id/submissions", submissionHandler.FindSubmissionByContestID)
			contest.GET("/:id/submissions/:submissionId", submissionHandler.FindByID)
		}
		v2.GET("/submissions/tasks", submissionHandler.GetTask)
		v2.POST("/submissions/tasks", submissionHandler.CreateSubmissionResult)
	}
}

func dummyHandler(c echo.Context) error {
	return c.String(400, "未実装のエンドポイントです")
}

func topHandler(c echo.Context) error {
	message := struct {
		Message string
	}{
		Message: "-----WELCOME TO Kemomimi Online Judge System v6-----\napi base url is '/api/v2'. learn more: github.com/poporonnet/kojs6-api-spec",
	}

	return c.JSON(http.StatusOK, message)
}
