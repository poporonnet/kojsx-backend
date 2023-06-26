package postgres

import (
	"database/sql"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/postgres/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestRepository struct {
	db *sql.DB
}

func NewContestRepository(db *sql.DB) *ContestRepository {
	return &ContestRepository{db: db}
}

func (c ContestRepository) CreateContest(d domain.Contest) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`INSERT INTO contests(id, title, description, startat, endat) VALUES ($1, $2, $3, $4, $5)`,
		d.GetID(),
		d.GetTitle(),
		d.GetDescription(),
		d.GetStartAt(),
		d.GetEndAt(),
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

func (c ContestRepository) FindAllContests() ([]domain.Contest, error) {
	rows, err := c.db.Query("SELECT * FROM contests;")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	res := make([]domain.Contest, 0)
	for rows.Next() {
		var e *entity.Contest
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.StartAt, &e.EndAt)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}
	return res, nil
}

func (c ContestRepository) FindContestByID(id id.SnowFlakeID) (*domain.Contest, error) {
	row := c.db.QueryRow(`SELECT * FROM contests WHERE id=?;`, id)
	var res *entity.Contest
	err := row.Scan(&res.ID, &res.Title, &res.Description, &res.StartAt, &res.EndAt)
	if err != nil {
		return nil, err
	}
	r := res.ToDomain()
	return &r, nil
}

func (c ContestRepository) FindContestByTitle(title string) (*domain.Contest, error) {
	row := c.db.QueryRow(`SELECT * FROM contests WHERE title=?;`, title)
	var res *entity.Contest
	err := row.Scan(&res.ID, &res.Title, &res.Description, &res.StartAt, &res.EndAt)
	if err != nil {
		return nil, err
	}
	r := res.ToDomain()
	return &r, nil
}
