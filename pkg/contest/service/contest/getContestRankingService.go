package contest

import (
	"errors"
	"sort"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	repository2 "github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	model2 "github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type GetContestRankingService struct {
	contestRepository    repository2.ContestRepository
	contestantRepository repository2.ContestantRepository
	problemRepository    repository2.ProblemRepository
	submissionRepository repository2.SubmissionRepository
	userRepository       repository.UserRepository
}

func NewGetContestRankingService(
	contest repository2.ContestRepository,
	contestant repository2.ContestantRepository,
	problem repository2.ProblemRepository,
	submission repository2.SubmissionRepository,
	user repository.UserRepository,
) *GetContestRankingService {
	return &GetContestRankingService{
		contestRepository:    contest,
		contestantRepository: contestant,
		problemRepository:    problem,
		submissionRepository: submission,
		userRepository:       user,
	}
}

// nolint
func (s GetContestRankingService) Handle(contestID id.SnowFlakeID) ([]UserFinalResult, error) {
	// コンテストの問題を取得
	problemsRes, err := s.problemRepository.FindProblemByContestID(contestID)
	if err != nil {
		utils.SugarLogger.Errorf("failed to get problems: %s", err)
		return nil, err
	}
	if len(problemsRes) == 0 {
		return nil, errors.New("not found")
	}
	// [ProblemID]domain.Problem
	problems := map[id.SnowFlakeID]model.Problem{}
	for _, v := range problemsRes {
		problems[v.GetProblemID()] = v
	}
	// コンテストの参加者を取得
	contestantsRes, err := s.contestantRepository.FindContestantByContestID(contestID)
	if err != nil {
		utils.SugarLogger.Errorf("failed to get contestants: %s", err)
		return nil, err
	}
	// [UserID]domain.Contestant
	contestants := map[id.SnowFlakeID]model.Contestant{}
	for _, v := range contestantsRes {
		if v.IsAdmin() || v.IsTester() {
			continue
		}
		contestants[v.GetUserID()] = v
	}
	// コンテストの参加者のユーザー情報を取得
	// [UserID]domain.User
	users := map[id.SnowFlakeID]model2.User{}
	for k := range contestants {
		user, err := s.userRepository.FindUserByID(k)
		if err != nil {
			utils.SugarLogger.Errorf("failed to get users: %s", err)
			return nil, err
		}
		users[user.GetID()] = *user
	}
	// 各問題に対する提出を取得
	submissionsRes := make([]model.Submission, 0)
	for _, v := range problems {
		submissionInProblem, err := s.submissionRepository.FindSubmissionByProblemID(v.GetProblemID())
		if err != nil {
			utils.SugarLogger.Errorf("failed to get submission: %s", err)
			return nil, err
		}
		submissionsRes = append(submissionsRes, submissionInProblem...)
	}
	// [ContestantID][]domain.Submission
	submissions := map[id.SnowFlakeID][]model.Submission{}
	for _, v := range submissionsRes {
		submissions[v.GetContestantID()] = append(submissions[v.GetContestantID()], v)
	}

	// 各ユーザーごとに結果を一時的に詰める
	results := make([]UserFinalResult, len(users))
	ii := 0
	for _, v := range contestants {
		// 管理者とテスターは載せない
		if v.IsAdmin() || v.IsTester() {
			continue
		}
		results[ii] = UserFinalResult{
			Point:       v.GetPoint(),
			User:        users[v.GetUserID()],
			Submissions: submissions[v.GetID()],
			Contestant:  v,
		}
		ii++
	}

	// 各ユーザーごとの結果から、得点を再計算する
	for i, v := range results {
		// [ProblemID]maxPoint 問題ごとの最高点数
		maxPoint := map[id.SnowFlakeID]int{}
		for _, j := range v.Submissions {
			if j.GetPoint() > maxPoint[j.GetProblemID()] {
				maxPoint[j.GetProblemID()] = j.GetPoint()
			}
		}
		point := 0
		for _, o := range maxPoint {
			point += o
		}

		results[i].Point = point
	}

	// 得点が高い順にソートする
	sort.Slice(results, func(i, j int) bool {
		return results[i].Point > results[j].Point
	})

	// ソートされた順に順位をつける
	rank := 0
	beforePoint := 0
	// 同じ点数だった場合は順位を変えない
	for i, v := range results {
		if v.Point != beforePoint || beforePoint == 0 {
			rank++
			beforePoint = v.Point
		}
		results[i].Rank = rank
	}

	return results, nil
}

type UserFinalResult struct {
	Rank        int
	Point       int
	User        model2.User
	Contestant  model.Contestant
	Submissions []model.Submission
}

/*
ランキング:
欲しい情報

提出ユーザー名
各問題に対する獲得点数
全体で取得した総点数
順位

*/
