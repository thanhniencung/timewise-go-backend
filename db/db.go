package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type SQL struct {
	Db *sqlx.DB
}

// Connect to postgres database
func (s *SQL) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"ryan",
		"postgres",
		"timewise")

	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection db ok")
}

// Close database connection
func (s *SQL) Close() {
	s.Db.Close()
}