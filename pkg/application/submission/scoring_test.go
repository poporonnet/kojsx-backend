package submission

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	prob domain.Problem
	rs   []domain.SubmissionResult
)

func init() {
	//　ケースを作る
	case1 := domain.NewCase("1", "1")
	_ = case1.SetIn("hello\n")
	_ = case1.SetOut("world\n")
	case2 := domain.NewCase("2", "2")
	_ = case2.SetIn("2\n")
	_ = case2.SetOut("4\n")

	// ケースセットを作る
	set1 := domain.NewCaseset("1")
	_ = set1.SetName("test 1")
	_ = set1.SetPoint(100)
	_ = set1.AddCase(*case1)
	set2 := domain.NewCaseset("2")
	_ = set2.SetName("test 2")
	_ = set2.SetPoint(200)
	_ = set2.AddCase(*case2)

	// 問題を作る
	p := domain.NewProblem("1", "1")
	_ = p.SetIndex("A")
	_ = p.SetTitle("test problem")
	_ = p.SetText("hello world")
	_ = p.SetPoint(300)
	_ = p.SetTimeLimit(2000)
	_ = p.AddCaseSet(*set1)
	_ = p.AddCaseSet(*set2)
	prob = *p

	rs = []domain.SubmissionResult{
		*domain.NewSubmissionResult("1", "WJ", "world\n", "1", 0, 300, 100),         // AC
		*domain.NewSubmissionResult("2", "WJ", "世界\n", "1", 0, 300, 100),            // WA
		*domain.NewSubmissionResult("3", "WJ", "", "1", 256, 10, 100),               // RE
		*domain.NewSubmissionResult("4", "WJ", "world\n", "1", 0, 3000, 100),        // TLE
		*domain.NewSubmissionResult("5", "WJ", "world\n", "1", 0, 300, 10000000000), // MLE
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
	filter := [][]domain.SubmissionResult{
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
