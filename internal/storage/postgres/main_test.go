package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	pool, err = NewDB("postgres://postgres:password@localhost:5432/prefood")
	if err != nil {
		log.Fatalf("failed to init storage: %s", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
