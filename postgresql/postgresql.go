package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type PGDatabase struct {
	db *pgxpool.Pool
}

var (
	pgdb   *PGDatabase
	dbOnce sync.Once
)

func GetPool(ctx context.Context, connString string) (*PGDatabase, error) {
	dbOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)

		if err != nil {
			fmt.Println("Cannot connect to DB:", err)
			return
		}

		pgdb = &PGDatabase{db}
	})

	return pgdb, nil
}

func (conn *PGDatabase) Ping(ctx context.Context) error {
	return conn.db.Ping(ctx)
}

func (conn *PGDatabase) Close() {
	conn.db.Close()
}
