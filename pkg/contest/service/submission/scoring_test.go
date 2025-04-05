package submission

import (
	"strconv"
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/stretchr/testify/assert"
)

var (
	prob model.Problem
	rs   []model.SubmissionResult
)

func init() {
	//　ケースを作る
	case1 := model.NewCase("1", "1")
	_ = case1.SetIn("hello\n")
	_ = case1.SetOut("world\n")
	case2 := model.NewCase("2", "2")
	_ = case2.SetIn("2\n")
	_ = case2.SetOut("4\n")

	// ケースセットを作る
	set1 := model.NewCaseset("1")
	_ = set1.SetName("test 1")
	_ = set1.SetPoint(100)
	_ = set1.AddCase(*case1)
	set2 := model.NewCaseset("2")
	_ = set2.SetName("test 2")
	_ = set2.SetPoint(200)
	_ = set2.AddCase(*case2)

	// 問題を作る
	p := model.NewProblem("1", "1")
	_ = p.SetIndex("A")
	_ = p.SetTitle("test problem")
	_ = p.SetText("hello world")
	_ = p.SetTimeLimit(2000)
	_ = p.AddCaseSet(*set1)
	_ = p.AddCaseSet(*set2)
	prob = *p

	rs = []model.SubmissionResult{
		*model.NewSubmissionResult("1", "WJ", "world\n", "1", 0, 300, 100),         // AC
		*model.NewSubmissionResult("2", "WJ", "世界\n", "1", 0, 300, 100),            // WA
		*model.NewSubmissionResult("3", "WJ", "", "1", 256, 10, 100),               // RE
		*model.NewSubmissionResult("4", "WJ", "world\n", "1", 0, 3000, 100),        // TLE
		*model.NewSubmissionResult("5", "WJ", "world\n", "1", 0, 300, 10000000000), // MLE
	}
}

func Test_judge(t *testing.T) {
	exp := []string{"AC", "WA", "RE", "TLE", "MLE"}
	for i, tt := range rs {
		t.Run(string(tt.GetID()), func(t *testing.T) {
			assert.Equal(t, exp[i], judge("world\n", tt, prob.GetTimeLimit(), prob.GetMemoryLimit()))
		})
	}
}

func Test_scoring(t *testing.T) {
	filter := [][]model.SubmissionResult{
		rs[0:1], // AC -> 100
		rs[:],   // MLE -> 0
		rs[1:2], // WA -> 0
		rs[0:3], // RE -> 0
		rs[0:4], // TLE -> 0
	}
	exp := []int{100, 0, 0, 0, 0}
	expStatus := []string{"AC", "MLE", "WA", "RE", "TLE"}
	for i := range filter {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			res, _ := scoring(prob, filter[i])
			assert.Equal(t, exp[i], res.Point)
			assert.Equal(t, expStatus[i], res.Status)
		})
	}
}
