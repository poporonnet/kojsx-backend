package domain

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
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
	*/
	if c.role == 1 {
		return true
	}
	return false
}

func (c *Contestant) GetPoint() int {
	return c.point
}

func (c *Contestant) SetAdmin() {
	if c.role == 0 {
		c.role = 1
	}
}

func (c *Contestant) SetNormal() {
	if c.role == 1 {
		c.role = 0
	}
}

func (c *Contestant) SetPoint(point int) error {
	if point%100 != 0 || point > 0 || point < 5000 {
		return errors.New("InvalidPointError")
	}

	c.point = point
	return nil
}
