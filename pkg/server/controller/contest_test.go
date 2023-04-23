package controller

import (
	"testing"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

var controller *ContestController

func init() {
	r := inmemory.NewContestRepository(dummyData.ContestArray)
	s := contest.NewCreateContestService(r)
	controller = NewContestController(r, *s)
}

func TestContestController_CreateContest(t *testing.T) {
	req := model.CreateContestRequestJSON{
		Title:       "Test Contest",
		Description: "# Test Contest\\nWelcome to Test Contest!\\n",
		StartAt:     time.Now(),
		EndAt:       time.Now().Add(1 * time.Hour),
	}

	res, _ := controller.CreateContest(req)

	assert.Equal(t, req.Title, res.Title)
	assert.Equal(t, req.Description, res.Description)
	assert.Equal(t, req.StartAt, res.StartAt)
	assert.Equal(t, req.EndAt, res.EndAt)
}
