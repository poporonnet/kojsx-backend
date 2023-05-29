package problem

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Data struct {
	id          id.SnowFlakeID
	contestID   id.SnowFlakeID
	index       string
	title       string
	text        string
	point       int
	memoryLimit int
	timeLimit   int
}

func (d Data) GetID() id.SnowFlakeID {
	return d.id
}

func (d Data) GetContestID() id.SnowFlakeID {
	return d.contestID
}

func (d Data) GetIndex() string {
	return d.index
}

func (d Data) GetTitle() string {
	return d.title
}

func (d Data) GetText() string {
	return d.text
}

func (d Data) GetPoint() int {
	return d.point
}

func (d Data) GetMemoryLimit() int {
	return d.memoryLimit
}

func (d Data) GetTimeLimit() int {
	return d.timeLimit
}

func NewData(
	id, contestID id.SnowFlakeID,
	index, title, text string,
	point, timeLimit int,
) *Data {
	return &Data{
		id:          id,
		contestID:   contestID,
		index:       index,
		title:       title,
		text:        text,
		point:       point,
		memoryLimit: domain.MemoryLimit,
		timeLimit:   timeLimit,
	}
}

func DataToDomain(in Data) domain.Problem {
	d := domain.NewProblem(
		in.GetID(),
		in.GetContestID(),
	)
	d.SetIndex(in.GetIndex())
	d.SetTitle(in.GetTitle())
	d.SetText(in.GetText())
	d.SetPoint(in.GetPoint())
	d.SetTimeLimit(in.GetTimeLimit())

	return *d
}

func DomainToData(in domain.Problem) Data {
	return *NewData(
		in.GetProblemID(),
		in.GetContestID(),
		in.GetIndex(),
		in.GetTitle(),
		in.GetText(),
		in.GetPoint(),
		in.GetMemoryLimit(),
	)
}
