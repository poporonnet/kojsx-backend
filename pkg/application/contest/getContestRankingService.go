package contest

import (
	"errors"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"sort"
)

type GetContestRankingService struct {
	contestRepository    repository.ContestRepository
	contestantRepository repository.ContestantRepository
	problemRepository    repository.ProblemRepository
	submissionRepository repository.SubmissionRepository
	userRepository       repository.UserRepository
}

func NewGetContestRankingService(
	contest repository.ContestRepository,
	contestant repository.ContestantRepository,
	problem repository.ProblemRepository,
	submission repository.SubmissionRepository,
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
	problems := map[id.SnowFlakeID]domain.Problem{}
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
	contestants := map[id.SnowFlakeID]domain.Contestant{}
	for _, v := range contestantsRes {
		if v.IsAdmin() || v.IsTester() {
			continue
		}
		contestants[v.GetUserID()] = v
	}
	// コンテストの参加者のユーザー情報を取得
	// [UserID]domain.User
	users := map[id.SnowFlakeID]domain.User{}
	for k := range contestants {
		user, err := s.userRepository.FindUserByID(k)
		if err != nil {
			utils.SugarLogger.Errorf("failed to get users: %s", err)
			return nil, err
		}
		users[user.GetID()] = *user
	}
	// 各問題に対する提出を取得
	submissionsRes := make([]domain.Submission, 0)
	for _, v := range problems {
		submissionInProblem, err := s.submissionRepository.FindSubmissionByProblemID(v.GetProblemID())
		if err != nil {
			utils.SugarLogger.Errorf("failed to get submission: %s", err)
			return nil, err
		}
		submissionsRes = append(submissionsRes, submissionInProblem...)
	}
	// [ContestantID][]domain.Submission
	submissions := map[id.SnowFlakeID][]domain.Submission{}
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
	User        domain.User
	Contestant  domain.Contestant
	Submissions []domain.Submission
}

/*
ランキング:
欲しい情報

提出ユーザー名
各問題に対する獲得点数
全体で取得した総点数
順位

*/
