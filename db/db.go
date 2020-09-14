package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"timewise/model"
)

type SQL struct {
	Db *sqlx.DB
}

// Connect to postgres database
func (s *SQL) Connect(cfg model.Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbUserName,
		cfg.Database.DbPassword,
		cfg.Database.DbName)

	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect database successfully")
}

// Close database connection
func (s *SQL) Close() {
	s.Db.Close()
}
