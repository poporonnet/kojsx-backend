package postgres

import (
	"database/sql"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/postgres/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestantRepository struct {
	db *sql.DB
}

func NewContestantRepository(db *sql.DB) *ContestantRepository {
	return &ContestantRepository{db: db}
}

func (c ContestantRepository) JoinContest(d domain.Contestant) error {
	role := domain.ContestParticipants
	if d.IsAdmin() {
		role = domain.ContestAdmin
	}
	if d.IsTester() {
		role = domain.ContestTester
	}

	e := entity.Contestant{
		ID:        string(d.GetID()),
		Role:      role,
		Point:     d.GetPoint(),
		ContestID: string(d.GetContestID()),
		UserID:    string(d.GetUserID()),
	}
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`INSERT INTO contestants(id, role, point, contestid, userid) VALUES ($1,$2,$3,$4,$5)`,
		e.ID,
		e.Role,
		e.Point,
		e.ContestID,
		e.UserID,
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

func (c ContestantRepository) FindContestantByID(id id.SnowFlakeID) (*domain.Contestant, error) {
	row := c.db.QueryRow(`SELECT * FROM contestants WHERE id=?;`, id)
	var res *entity.Contestant
	err := row.Scan(&res.ID, &res.Role, &res.Point, &res.ContestID, &res.UserID)
	if err != nil {
		return nil, err
	}

	r := res.ToDomain()
	return &r, nil
}

func (c ContestantRepository) FindContestantByUserID(id id.SnowFlakeID) ([]domain.Contestant, error) {
	rows, err := c.db.Query(`SELECT * FROM contestants WHERE userid=?;`, id)
	if err != nil {
		return nil, err
	}
	var res []domain.Contestant
	for rows.Next() {
		e := &entity.Contestant{}
		err := rows.Scan(&e.ID, &e.Role, &e.Point, &e.ContestID, &e.UserID)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}
	return res, nil
}

func (c ContestantRepository) FindContestantByContestID(id id.SnowFlakeID) ([]domain.Contestant, error) {
	rows, err := c.db.Query(`SELECT * FROM contestants WHERE contestid=?;`, id)
	if err != nil {
		return nil, err
	}
	var res []domain.Contestant
	for rows.Next() {
		e := &entity.Contestant{}
		err := rows.Scan(&e.ID, &e.Role, &e.Point, &e.ContestID, &e.UserID)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}
	return res, nil
}
