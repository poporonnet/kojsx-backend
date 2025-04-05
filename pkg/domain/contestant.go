package domain

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestantRole int

const (
	ContestParticipants = iota
	ContestAdmin
	ContestTester
)

type Contestant struct {
	id        id.SnowFlakeID
	contestID id.SnowFlakeID
	userID    id.SnowFlakeID
	role      int
	point     int
}

// NewContestant 不変値: ID/ContestID/UserID
func NewContestant(id id.SnowFlakeID, cID id.SnowFlakeID, uID id.SnowFlakeID) *Contestant {
	return &Contestant{id: id, contestID: cID, userID: uID}
}

func (c *Contestant) GetID() id.SnowFlakeID {
	return c.id
}

func (c *Contestant) GetContestID() id.SnowFlakeID {
	return c.contestID
}

func (c *Contestant) GetUserID() id.SnowFlakeID {
	return c.userID
}

func (c *Contestant) IsAdmin() bool {
	/*
		ロール:
			0 参加者
			1 アドミン
			2 テスター
	*/
	if c.role == ContestAdmin {
		return true
	}
	return false
}

func (c *Contestant) IsTester() bool {
	return c.role == ContestTester
}

func (c *Contestant) IsNormal() bool {
	return c.role != ContestAdmin && c.role != ContestTester
}

func (c *Contestant) GetPoint() int {
	return c.point
}

func (c *Contestant) SetAdmin() {
	if c.role == ContestParticipants {
		c.role = ContestAdmin
	}
}

func (c *Contestant) SetNormal() {
	if c.role != ContestParticipants {
		c.role = ContestParticipants
	}
}

func (c *Contestant) SetTester() {
	if c.role != ContestTester {
		c.role = ContestTester
	}
}

func (c *Contestant) SetPoint(point int) error {
	if point%100 != 0 || point > 0 || point < 5000 {
		return errors.New("InvalidPointError")
	}

	c.point = point
	return nil
}
