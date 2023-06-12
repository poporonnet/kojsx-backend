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

	caseSets []CaseSetData
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

func (d Data) GetCaseSets() []CaseSetData {
	return d.caseSets
}

func NewData(
	id, contestID id.SnowFlakeID,
	index, title, text string,
	point, timeLimit int,
	sets []CaseSetData,
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
		caseSets:    sets,
	}
}

func (d Data) ToDomain() *domain.Problem {
	dd := domain.NewProblem(
		d.GetID(),
		d.GetContestID(),
	)
	err := dd.SetIndex(d.GetIndex())
	if err != nil {
		return nil
	}
	err = dd.SetTitle(d.GetTitle())
	if err != nil {
		return nil
	}
	err = dd.SetText(d.GetText())
	if err != nil {
		return nil
	}
	for _, v := range d.GetCaseSets() {
		_ = dd.AddCaseSet(*v.ToDomain())
	}
	err = dd.SetTimeLimit(d.GetTimeLimit())
	if err != nil {
		return nil
	}

	return dd
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
		DomainToCaseSetData(in.GetCaseSets()),
	)
}

type CaseSetData struct {
	id    id.SnowFlakeID
	name  string
	point int

	cases []CaseData
}

func (d CaseSetData) GetID() id.SnowFlakeID {
	return d.id
}

func (d CaseSetData) GetName() string {
	return d.name
}

func (d CaseSetData) GetPoint() int {
	return d.point
}

func (d CaseSetData) GetCases() []CaseData {
	return d.cases
}

func NewCaseSetData(id id.SnowFlakeID, name string, point int, cases []CaseData) *CaseSetData {
	return &CaseSetData{
		id:    id,
		name:  name,
		point: point,
		cases: cases,
	}
}

func (d CaseSetData) ToDomain() *domain.Caseset {
	s := domain.NewCaseset(d.id)
	_ = s.SetName(d.name)
	_ = s.SetPoint(d.point)

	for _, v := range d.GetCases() {
		_ = s.AddCase(*v.ToDomain())
	}

	return s
}

func DomainToCaseSetData(in []domain.Caseset) []CaseSetData {
	res := make([]CaseSetData, len(in))
	for i, v := range in {
		cases := make([]CaseData, len(v.GetCases()))
		for j, k := range v.GetCases() {
			cases[j] = *NewCaseData(
				k.GetID(),
				k.GetCasesetID(),
				k.GetIn(),
				k.GetOut(),
			)
		}

		res[i] = *NewCaseSetData(
			v.GetID(),
			v.GetName(),
			v.GetPoint(),
			cases,
		)
	}
	return res
}

type CaseData struct {
	id        id.SnowFlakeID
	caseSetID id.SnowFlakeID
	in        string
	out       string
}

func (d CaseData) GetID() id.SnowFlakeID {
	return d.id
}

func (d CaseData) GetCaseSetID() id.SnowFlakeID {
	return d.caseSetID
}

func (d CaseData) GetIn() string {
	return d.in
}

func (d CaseData) GetOut() string {
	return d.out
}

func NewCaseData(id, caseSetID id.SnowFlakeID, in, out string) *CaseData {
	return &CaseData{
		id:        id,
		caseSetID: caseSetID,
		in:        in,
		out:       out,
	}
}

func (d CaseData) ToDomain() *domain.Case {
	c := domain.NewCase(d.id, d.caseSetID)
	_ = c.SetIn(d.in)
	_ = c.SetOut(d.out)
	return c
}
