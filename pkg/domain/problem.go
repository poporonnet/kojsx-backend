package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Problem struct {
	id          id.SnowFlakeID
	contestID   id.SnowFlakeID
	index       string
	title       string
	text        string
	point       int
	memoryLimit int
	timeLimit   int

	caseSet []Caseset
}

// ProblemIndexInvalidError 問題の添字制約エラー
type ProblemIndexInvalidError struct {
}

func (e ProblemIndexInvalidError) Error() string {
	return "問題添字が不正です"
}

// ProblemTitleLengthError 問題のタイトル制約エラー
type ProblemTitleLengthError struct {
}

func (e ProblemTitleLengthError) Error() string {
	return "問題のタイトルが不正です"
}

// ProblemTextLengthError 問題文の長さ制約エラー
type ProblemTextLengthError struct {
}

func (p ProblemTextLengthError) Error() string {
	return "問題文の長さが不正です"
}

// ProblemPointInvalidError 問題ポイント制約エラー
type ProblemPointInvalidError struct {
}

func (ProblemPointInvalidError) Error() string {
	return "問題の得点が不正です"
}

// ProblemTimeLimitInvalidError 問題実行時間制約エラー
type ProblemTimeLimitInvalidError struct {
}

func (ProblemTimeLimitInvalidError) Error() string {
	return "問題の実行時間が不正です"
}

// MemoryLimit 512 MB = 512000 KB
const MemoryLimit = 512000

/*
NewProblem
不変値: ID, ContestID
定数: MemoryLimit
*/
func NewProblem(pID id.SnowFlakeID, contestID id.SnowFlakeID) *Problem {
	return &Problem{
		id:          pID,
		contestID:   contestID,
		memoryLimit: MemoryLimit,
	}
}

func (p *Problem) GetProblemID() id.SnowFlakeID {
	return p.id
}

func (p *Problem) GetContestID() id.SnowFlakeID {
	return p.contestID
}

func (p *Problem) GetIndex() string {
	return p.index
}

func (p *Problem) GetTitle() string {
	return p.title
}

func (p *Problem) GetText() string {
	return p.text
}

func (p *Problem) GetPoint() int {
	return p.point
}

func (p *Problem) GetMemoryLimit() int {
	return p.memoryLimit
}

func (p *Problem) GetTimeLimit() int {
	return p.timeLimit
}

func (p *Problem) AddCaseSet(in Caseset) error {
	if len(p.caseSet) > 64 {
		return errors.New("これ以上ケースセットを追加できません")
	}

	if (p.point+in.GetPoint()) < 0 || (p.point+in.GetPoint()) > 5000 || (p.point+in.GetPoint())%100 != 0 {
		return ProblemPointInvalidError{}
	}
	p.caseSet = append(p.caseSet, in)
	// 得点はケースセット追加時に自動で計算される
	p.point += in.GetPoint()
	return nil
}

func (p *Problem) GetCaseSets() []Caseset {
	return p.caseSet
}

func (p *Problem) SetIndex(index string) error {
	// Indexの制約: アルファベット大文字、1文字以上3文字以下(コンテスト最大問題数が64なため
	if utf8.RuneCountInString(index) > 2 {
		return ProblemIndexInvalidError{}
	}
	p.index = index
	return nil
}

func (p *Problem) SetTitle(title string) error {
	// 64文字以下
	if utf8.RuneCountInString(title) > 64 {
		return ProblemTitleLengthError{}
	}
	p.title = title
	return nil
}

func (p *Problem) SetText(text string) error {
	// 50000文字以下
	if utf8.RuneCountInString(text) > 50000 {
		return ProblemTextLengthError{}
	}
	p.text = text
	return nil
}

func (p *Problem) SetTimeLimit(limit int) error {
	// 1~2000の間(単位ms) 10刻み
	if limit < 1 || limit > 2000 || limit%10 != 0 {
		return ProblemTimeLimitInvalidError{}
	}
	p.timeLimit = limit
	return nil
}
