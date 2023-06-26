package postgres

import (
	"database/sql"
	"fmt"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/postgres/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type SubmissionRepository struct {
	db *sql.DB
}

func NewSubmissionRepository(db *sql.DB) *SubmissionRepository {
	return &SubmissionRepository{
		db: db,
	}
}

func (s SubmissionRepository) CreateSubmission(submission domain.Submission) error {
	e := entity.Submission{
		ID:           string(submission.GetID()),
		Point:        submission.GetPoint(),
		Lang:         submission.GetLang(),
		CodeLength:   submission.GetCodeLength(),
		Result:       submission.GetResult(),
		ExecTime:     submission.GetExecTime(),
		ExecMemory:   submission.GetExecMemory(),
		Code:         submission.GetCode(),
		SubmittedAt:  submission.GetSubmittedAt(),
		ProblemID:    string(submission.GetProblemID()),
		ContestantID: string(submission.GetContestantID()),
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	fmt.Println(e.ProblemID,
		e.ContestantID)
	_, err = tx.Exec(
		`INSERT INTO submissions(id, point, lang, codelength, result, exectime, execmemory, code, submittedat, problemid, contestantid) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		e.ID,
		e.Point,
		e.Lang,
		e.CodeLength,
		e.Result,
		e.ExecTime,
		e.ExecMemory,
		e.Code,
		e.SubmittedAt,
		e.ProblemID,
		e.ContestantID,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (s SubmissionRepository) FindSubmissionByID(id id.SnowFlakeID) (*domain.Submission, error) {
	row := s.db.QueryRow(`SELECT * FROM submissions WHERE id=$1`, id)
	var resp entity.Submission
	err := row.Scan(&resp.ID, &resp.Point, &resp.Lang, &resp.CodeLength, &resp.Result, &resp.ExecTime, &resp.ExecMemory, &resp.Code, &resp.SubmittedAt, &resp.ProblemID, &resp.ContestantID)
	if err != nil {
		return nil, err
	}
	res := resp.ToDomain()

	return &res, nil
}

func (s SubmissionRepository) FindSubmissionByStatus(status string) ([]domain.Submission, error) {
	rows, err := s.db.Query(`SELECT * FROM submissions WHERE result=$1;`, status)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	res := make([]domain.Submission, 0)
	for rows.Next() {
		var e entity.Submission
		err := rows.Scan(&e.ID, &e.Point, &e.Lang, &e.CodeLength, &e.Result, &e.ExecTime, &e.ExecMemory, &e.Code, &e.SubmittedAt, &e.ProblemID, &e.ContestantID)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}

	return res, nil
}

func (s SubmissionRepository) FindSubmissionByProblemID(id id.SnowFlakeID) ([]domain.Submission, error) {
	rows, err := s.db.Query(`SELECT * FROM submissions WHERE id=$1;`, id)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	res := make([]domain.Submission, 0)
	for rows.Next() {
		var e entity.Submission
		err := rows.Scan(&e.ID, &e.Point, &e.Lang, &e.CodeLength, &e.Result, &e.ExecTime, &e.ExecMemory, &e.Code, &e.SubmittedAt, &e.ProblemID, &e.ContestantID)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}

	return res, nil
}

func (s SubmissionRepository) UpdateSubmissionResult(submission domain.Submission) (*domain.Submission, error) {
	//TODO implement me
	panic("implement me")
}
