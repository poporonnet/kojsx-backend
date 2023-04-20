package contest

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestFindContestService_FindByID(t *testing.T) {
	r := inmemory.NewContestRepository(dummyData.ContestArray)
	s := NewFindContestService(r)

	// 取得できる
	r1, _ := s.FindByID("1")
	assert.Equal(t, DomainToData(*dummyData.ExistsContestData), *r1)
	// 取得できない
	r2, _ := s.FindByID("9")
	var ex *Data = nil
	assert.Equal(t, ex, r2)
}

func TestFindContestService_FindAll(t *testing.T) {
	r := inmemory.NewContestRepository(dummyData.ContestArray)
	s := NewFindContestService(r)

	// 取得できる
	r1, _ := s.FindAll()
	ex := make([]Data, len(dummyData.ContestArray))
	for i, v := range dummyData.ContestArray {
		ex[i] = DomainToData(v)
	}

	assert.Equal(t, ex, r1)
}
