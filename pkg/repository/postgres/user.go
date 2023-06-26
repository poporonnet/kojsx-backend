package postgres

import (
	"database/sql"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/postgres/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) CreateUser(d domain.User) error {
	role := domain.Unverified
	if d.IsVerified() {
		role = domain.Normal
	}
	if d.IsAdmin() {
		role = domain.Admin
	}

	e := entity.User{
		ID:       string(d.GetID()),
		Name:     d.GetName(),
		Email:    d.GetEmail(),
		Password: d.GetPassword(),
		Role:     role,
	}

	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`INSERT INTO "users"(id, name, email, password, role) VALUES ($1, $2, $3, $4, $5);`,
		e.ID,
		e.Name,
		e.Email,
		e.Password,
		e.Role,
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

func (u UserRepository) FindAllUsers() ([]domain.User, error) {
	rows, err := u.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	res := make([]domain.User, 0)
	for rows.Next() {
		var e *entity.User
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Password, &e.Role)
		if err != nil {
			return nil, err
		}
		res = append(res, e.ToDomain())
	}
	return res, nil
}

func (u UserRepository) FindUserByID(id id.SnowFlakeID) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT * FROM users WHERE id=?;`, id)
	var res *entity.User
	err := row.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Role)
	if err != nil {
		return nil, err
	}
	r := res.ToDomain()
	return &r, nil
}

func (u UserRepository) FindUserByName(name string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT * FROM users WHERE name=?;`, name)
	var res *entity.User
	err := row.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Role)
	if err != nil {
		return nil, err
	}
	r := res.ToDomain()
	return &r, nil
}

func (u UserRepository) FindUserByEmail(email string) (*domain.User, error) {
	row := u.db.QueryRow(`SELECT * FROM users WHERE email=?;`, email)
	var res *entity.User
	err := row.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Role)
	if err != nil {
		return nil, err
	}
	r := res.ToDomain()
	return &r, nil
}

func (u UserRepository) UpdateUser(d domain.User) error {
	//TODO implement me
	panic("implement me")
}
