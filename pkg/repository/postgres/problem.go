package postgres

import (
	"database/sql"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/postgres/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ProblemRepository struct {
	db *sql.DB
}

func NewProblemRepository(db *sql.DB) *ProblemRepository {
	return &ProblemRepository{db: db}
}

func (p ProblemRepository) CreateProblem(in domain.Problem) error {
	e := entity.Problem{
		ID:          string(in.GetProblemID()),
		Index:       in.GetIndex(),
		Title:       in.GetTitle(),
		Text:        in.GetText(),
		Point:       in.GetPoint(),
		MemoryLimit: in.GetMemoryLimit(),
		TimeLimit:   in.GetTimeLimit(),
		ContestID:   string(in.GetContestID()),
	}

	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`INSERT INTO problems(id, index, title, text, point, memorylimit, timelimit, contestid) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		e.ID,
		e.Index,
		e.Title,
		e.Text,
		e.Point,
		e.MemoryLimit,
		e.TimeLimit,
		e.ContestID,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// ケース/ケースセットを作る
	// すべての作成が完了しなければトランザクションは成功しない
	for _, v := range in.GetCaseSets() {
		err = p.addCaseSet(tx, v, in.GetProblemID())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

// ケースセットを作成
func (p ProblemRepository) addCaseSet(tx *sql.Tx, d domain.Caseset, id id.SnowFlakeID) error {
	e := entity.CaseSet{
		ID:        string(d.GetID()),
		Name:      d.GetName(),
		Point:     d.GetPoint(),
		ProblemID: string(id),
	}

	_, err := tx.Exec(
		`INSERT INTO casesets(id, name, point, problemid) VALUES ($1,$2,$3,$4)`,
		e.ID,
		e.Name,
		e.Point,
		e.ProblemID,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, v := range d.GetCases() {
		err = p.addCase(tx, v)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return nil
}

// ケースを作成
func (p ProblemRepository) addCase(tx *sql.Tx, d domain.Case) error {
	e := entity.Case{
		ID:        string(d.GetID()),
		Input:     d.GetIn(),
		Output:    d.GetOut(),
		CasesetID: string(d.GetCasesetID()),
	}

	_, err := tx.Exec(
		`INSERT INTO cases(id, input, output, casesetid) VALUES ($1,$2,$3,$4)`,
		e.ID,
		e.Input,
		e.Output,
		e.CasesetID,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

func (p ProblemRepository) FindProblemByContestID(id id.SnowFlakeID) ([]domain.Problem, error) {
	// 問題を取得
	rows, err := p.db.Query(`SELECT * FROM problems WHERE contestid=?`, id)
	if err != nil {
		return nil, err
	}

	var res []domain.Problem
	for rows.Next() {
		e := &entity.Problem{}
		err := rows.Scan(
			&e.ID,
			&e.Index,
			&e.Title,
			&e.Text,
			&e.Point,
			&e.MemoryLimit,
			&e.TimeLimit,
			&e.ContestID,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}

	return res, err
}

func (p ProblemRepository) FindProblemByID(id id.SnowFlakeID) (*domain.Problem, error) {
	row := p.db.QueryRow(`SELECT * FROM problems WHERE id=?`, id)

	var res *entity.Problem
	err := row.Scan(
		&res.ID,
		&res.Index,
		&res.Title,
		&res.Text,
		&res.Point,
		&res.MemoryLimit,
		&res.TimeLimit,
		&res.ContestID,
	)
	if err != nil {
		return nil, err
	}

	r := res.ToDomain()
	return &r, nil
}

func (p ProblemRepository) FindProblemByTitle(name string) (*domain.Problem, error) {
	row := p.db.QueryRow(`SELECT * FROM problems WHERE title=?`, name)
	var res *entity.Problem
	err := row.Scan(
		&res.ID,
		&res.Index,
		&res.Title,
		&res.Text,
		&res.Point,
		&res.MemoryLimit,
		&res.TimeLimit,
		&res.ContestID,
	)
	if err != nil {
		return nil, err
	}

	r := res.ToDomain()
	return &r, nil
}

func (p ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) (*domain.Caseset, error) {
	row := p.db.QueryRow(`SELECT * FROM casesets WHERE id=?`, id)
	var res *entity.CaseSet
	err := row.Scan(
		&res.ID,
		&res.Name,
		&res.Point,
		&res.ProblemID,
	)
	if err != nil {
		return nil, err
	}

	r := res.ToDomain()
	return &r, nil
}

func (p ProblemRepository) FindCaseByID(id id.SnowFlakeID) (*domain.Case, error) {
	row := p.db.QueryRow(`SELECT * FROM cases WHERE id=?`, id)
	var res *entity.Case
	err := row.Scan(
		&res.ID,
		&res.Input,
		&res.Output,
		&res.CasesetID,
	)
	if err != nil {
		return nil, err
	}

	r := res.ToDomain()
	return &r, nil
}
