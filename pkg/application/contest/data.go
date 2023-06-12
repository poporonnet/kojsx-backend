package contest

import (
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Data struct {
	id          id.SnowFlakeID
	title       string
	description string
	startAt     time.Time
	endAt       time.Time
}

func NewData(id id.SnowFlakeID, title string, description string, startAt time.Time, endAt time.Time) *Data {
	return &Data{id: id, title: title, description: description, startAt: startAt, endAt: endAt}
}

func (d Data) GetID() id.SnowFlakeID {
	return d.id
}

func (d Data) GetTitle() string {
	return d.title
}

func (d Data) GetDescription() string {
	return d.description
}

func (d Data) GetStartAt() time.Time {
	return d.startAt
}

func (d Data) GetEndAt() time.Time {
	return d.endAt
}

// ToDomain DTOをドメインモデルに
func (d Data) ToDomain() domain.Contest {
	u := domain.NewContest(d.GetID())
	_ = u.SetTitle(d.GetTitle())
	_ = u.SetDescription(d.GetDescription())
	_ = u.SetStartAt(d.GetStartAt())
	_ = u.SetEndAt(d.GetEndAt())
	return *u
}

// DomainToData ドメインモデルをDTOに
func DomainToData(in domain.Contest) Data {
	return *NewData(in.GetID(), in.GetTitle(), in.GetDescription(), in.GetStartAt(), in.GetEndAt())
}
