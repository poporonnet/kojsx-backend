package contest

import (
	"testing"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestCreateContestService_Handle(t *testing.T) {
	r := inmemory.NewContestRepository(dummyData.ContestArray)
	cContestService := NewCreateContestService(r)

	// 成功するとき
	_, err := cContestService.Handle("Test Contest", "hello world", time.Now(), time.Now().Add(60*time.Minute))
	assert.Equal(t, nil, err)

	// 失敗するとき
	// 重複する
	_, err2 := cContestService.Handle("Test Contest", "hello", time.Now(), time.Now().Add(60*time.Minute))
	assert.NotEqual(t, nil, err2)
}
