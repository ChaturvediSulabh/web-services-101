package database

import (
	"database/sql"
	"log"
	"os"
)

//DbConn ...
var DbConn *sql.DB

//SetupDB ...
func SetupDB() *sql.DB {
	DbConn, err := sql.Open("postgres", os.Getenv("CONN_STR"))
	if err != nil {
		log.Fatal(err)
	}
	return DbConn
}

//GetAllData ...
func GetAllData(DbConn *sql.DB) []string {
	query := `SELECT item FROM "public"."menu"`
	rows, err := DbConn.Query(query)
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
