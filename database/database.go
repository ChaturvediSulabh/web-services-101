package database

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"time"
)

//DbConn ...
var DbConn *sql.DB

//ConnStr ...
var ConnStr = flag.String("DB_CONN_STR", "", "POSTGRES: Database connection string")

//SetupDB ...
func SetupDB() *sql.DB {
	flag.Parse()
	DbConn, err := sql.Open("postgres", *ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetConnMaxLifetime(60 * time.Second)
	DbConn.SetMaxIdleConns(4)
	DbConn.SetMaxOpenConns(4)
	return DbConn
}

//GetAllData ...
func GetAllData(DbConn *sql.DB) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	query := `SELECT item FROM "public"."menu"`
	rows, err := DbConn.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result []string
	for rows.Next() {
		var item string
		err := rows.Scan(&item)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, item)
	}
	return result
}
