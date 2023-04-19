package user

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService_Handle(t *testing.T) {
	r := inmemory.NewUserRepository(dummyData.UserArray)
	uService := service.NewUserService(r)
	cUserService := NewCreateUserService(r, *uService)

	// 成功するとき
	err := cUserService.Handle("yamada", "yamada@example.jp")
	assert.Equal(t, nil, err)

	// 失敗するとき
	// 重複したとき
	err2 := cUserService.Handle("yamada", "yamada@example.jp")
	assert.NotEqual(t, nil, err2)
}
