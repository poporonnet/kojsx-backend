package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type Case struct {
	id        id.SnowFlakeID
	casesetID id.SnowFlakeID
	in        string
	out       string
}

func (c *Case) GetID() id.SnowFlakeID {
	return c.id
}

func (c *Case) GetCasesetID() id.SnowFlakeID {
	return c.casesetID
}

func (c *Case) GetIn() string {
	return c.in
}

func (c *Case) GetOut() string {
	return c.out
}

// NewCase 不変値: id/casesetID
func NewCase(id id.SnowFlakeID, casesetID id.SnowFlakeID) *Case {
	return &Case{id: id, casesetID: casesetID}
}

func (c *Case) SetIn(i string) error {
	// 0~5000文字
	if utf8.RuneCountInString(i) < 0 || utf8.RuneCountInString(i) > 5000 {
		return errors.New("CaseInLengthError")
	}
	c.in = i
	return nil
}

func (c *Case) SetOut(i string) error {
	if utf8.RuneCountInString(i) < 0 || utf8.RuneCountInString(i) > 5000 {
		return errors.New("CaseInLengthError")
	}
	c.out = i
	return nil
}
