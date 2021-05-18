package sql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) (err error) {
	s.db, err = sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("cannot connect db: %w", err)
	}

	err = s.db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("cannot ping context: %w", err)
	}

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
