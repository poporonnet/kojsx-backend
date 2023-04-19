package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	a, _ = domain.NewUser("1", "test", "test@example.com")
	b, _ = domain.NewUser("2", "taro", "taro@example.com")

	NotExists, _ = domain.NewUser("3", "Yamada", "yamada@example.jp")
	Exists       = *a

	// UserArray 重複しないデータ
	UserArray = []domain.User{*a, *b}
)
