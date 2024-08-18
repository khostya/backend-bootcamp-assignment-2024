//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/backend-bootcamp-assignment-2024/tests/postgres/postgresql"
	"os"
	"testing"
)

var (
	db *postgresql.DBPool
)

const (
	usersTable         = "bootcamp.users"
	flatsTable         = "bootcamp.flats"
	housesTable        = "bootcamp.houses"
	subscriptionsTable = "bootcamp.subscriptions"
)

func TestMain(m *testing.M) {
	db = postgresql.NewFromEnv()

	code := m.Run()

	db.TruncateTable(context.Background(), subscriptionsTable, flatsTable, housesTable, usersTable)
	db.Close()

	os.Exit(code)
}
