package submission

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ScoreResult struct {
	Point            int
	SubmissionResult []model.SubmissionResult
	Status           string
}

// scoring
func scoring(problem model.Problem, results []model.SubmissionResult) (ScoreResult, error) {
	// 1. ケースセットごとに提出結果をまとめる
	// まとめる用のmap
	setGroup := map[id.SnowFlakeID][]model.SubmissionResult{}

	// 全結果について調べる
	for _, v := range results {
		// 結果のケースセットを調べる
		set, err := findCaseSetByCaseID(problem, id.SnowFlakeID(v.GetCaseName()))
		if err != nil {
			return ScoreResult{}, err
		}

		// ケースセットごとにまとめる
		setGroup[set.GetID()] = append(setGroup[set.GetID()], v)
	}

	// 2. statusを求める
	// ケースごとのStatus [ケースセットID][]status
	caseStatuses := map[id.SnowFlakeID][]string{}
	// ケースセットごとのStatus [ケースセットID]ケースセット全体のStatus
	setStatues := map[id.SnowFlakeID]string{}
	res := make([]model.SubmissionResult, 0)
	// ケースごとのStatusを求める
	for k, v := range setGroup {
		// ケースセットごとの提出
		for _, j := range v {
			// ケースを持ってくる
			c, err := findCase(problem, id.SnowFlakeID(j.GetCaseName()))
			if err != nil {
				return ScoreResult{}, err
			}
			caseStatuses[k] = append(caseStatuses[j.GetID()], judge(c.GetOut(), j, problem.GetTimeLimit(), problem.GetMemoryLimit()))
			res = append(res, *model.NewSubmissionResult(
				j.GetID(),
				judge(c.GetOut(), j, problem.GetTimeLimit(), problem.GetMemoryLimit()),
				j.GetOutput(),
				j.GetCaseName(),
				j.GetExitStatus(),
				j.GetExecTime(),
				j.GetExecMemory(),
			))
		}
	}
	// セットごとのStatusを求める
	for j, v := range caseStatuses {
		for _, k := range v {
			if k != "AC" && k != "IE" {
				setStatues[j] = k
				continue
			}
			setStatues[j] = k
		}
	}
	// 3. ケースセットごとに得点を計算して、足す
	point := 0
	status := ""
	for k, v := range setStatues {
		// 不正解
		if v != "AC" && v != "IE" {
			// なにもしない
			point = 0
			status = v
			break
		}
		// 正解
		if v == "AC" {
			// 満点
			set, err := findCaseSetByID(problem, k)
			if err != nil {
				return ScoreResult{}, err
			}
			point += set.GetPoint()
			status = v
		} else {
			// エラー
			set, err := findCaseSetByCaseID(problem, k)
			if err != nil {
				return ScoreResult{}, err
			}
			point += set.GetPoint() / 10
			status = v
			break
		}
	}

	return ScoreResult{
		point,
		res,
		status,
	}, nil
}

func judge(out string, r model.SubmissionResult, timeLim, memLim int) string {
	// 終了コードが0でない -> RE
	if r.GetExitStatus() != 0 {
		return "RE"
	}
	// 実行時間が規定値より大きい -> TLE
	if r.GetExecTime() > timeLim {
		return "TLE"
	}
	// メモリ使用量が規定値より多い -> MLE
	if r.GetExecMemory() > memLim {
		return "MLE"
	}
	// CE (現状判定できない？)
	// 想定解と違う -> WA
	if r.GetOutput() != out {
		return "WA"
	}
	return "AC"
}

// findCase 問題からケースを取得
func findCase(in model.Problem, id id.SnowFlakeID) (model.Case, error) {
	for _, v := range in.GetCaseSets() {
		for _, k := range v.GetCases() {
			if k.GetID() == id {
				return k, nil
			}
		}
	}

	utils.Logger.Sugar().Errorf("failed to find case: not found")
	return model.Case{}, errors.New("not found")
}

// findCaseSetByCaseID ケースセットを取得
func findCaseSetByCaseID(in model.Problem, id id.SnowFlakeID) (model.Caseset, error) {
	for _, v := range in.GetCaseSets() {
		for _, k := range v.GetCases() {
			if k.GetID() == id {
				return v, nil
			}
		}
	}
	utils.Logger.Sugar().Errorf("failed to find caseSet: not found")
	return model.Caseset{}, errors.New("not found")
}

func findCaseSetByID(in model.Problem, id id.SnowFlakeID) (model.Caseset, error) {
	for _, v := range in.GetCaseSets() {
		if v.GetID() == id {
			return v, nil
		}
	}
	utils.Logger.Sugar().Errorf("failed to find caseSet: not found")
	return model.Caseset{}, errors.New("not found")
}
