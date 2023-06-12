package controller_test

import (
	"testing"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/stretchr/testify/assert"
)

var cont *controller.ContestController

func init() {
	r := inmemory.NewContestRepository(seed.NewSeeds().Contests)
	s := contest.NewCreateContestService(r)
	f := contest.NewFindContestService(r)
	cont = controller.NewContestController(r, *s, *f)
}

func TestContestController_CreateContest(t *testing.T) {
	req := model.CreateContestRequestJSON{
		Title:       "Test Contest",
		Description: "# Test Contest\\nWelcome to Test Contest!\\n",
		StartAt:     time.Now(),
		EndAt:       time.Now().Add(1 * time.Hour),
	}

	res, _ := cont.CreateContest(req)

	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
	assert.Equal(t, req.StartAt, res.StartAt)
	assert.Equal(t, req.EndAt, res.EndAt)
}

func TestContestController_FindContestByID(t *testing.T) {
	res, _ := cont.FindContestByID("10")
	e := seed.NewSeeds().Contests
	assert.Equal(t, e[0].GetTitle(), res.Title)
	assert.Equal(t, e[0].GetDescription(), res.Description)
	assert.Equal(t, e[0].GetStartAt(), res.StartAt)
	assert.Equal(t, e[0].GetEndAt(), res.EndAt)
}
