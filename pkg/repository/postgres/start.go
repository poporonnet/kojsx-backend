package postgres

import "database/sql"

func NewConnection() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost port=5432 user=laminne dbname=kojs sslmode=disable")
}
