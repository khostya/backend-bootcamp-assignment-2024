//go:build integration

package postgres

import (
	"backend-bootcamp-assignment-2024/tests/postgres/postgresql"
	"context"
	"os"
	"testing"
)

var (
	db *postgresql.DBPool
)

const (
	usersTable  = "bootcamp.users"
	flatsTable  = "bootcamp.users"
	housesTable = "bootcamp.houses"
)

func TestMain(m *testing.M) {
	db = postgresql.NewFromEnv()

	code := m.Run()

	db.TruncateTable(context.Background(), flatsTable, housesTable, usersTable)
	db.Close()

	os.Exit(code)
}
