package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	k = domain.NewCaseset("1")
	l = domain.NewCaseset("2")

	ExistsCasesetData    = k
	NotExistsCasesetData = domain.NewCaseset("3")

	CasesetArray = []domain.Caseset{*k, *l}
)
