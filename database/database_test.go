package database

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestSetupDB(t *testing.T) {
	// var ConnStr = flag.String("DB_CONN_STR", "", "POSTGRES: Database connection string")
	// flag.Parse()
	DbConn := SetupDB()
	query := `SELECT item FROM "public"."menu"`
	_, err := DbConn.Query(query)
	if err != nil {
		t.FailNow()
	}
}
