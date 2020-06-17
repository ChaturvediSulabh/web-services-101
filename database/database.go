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
func GetAllData(DbConn *sql.DB) (string, error) {
	query := `SELECT * FROM "public"."menu" LIMIT 100`
	rows, err := DbConn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var item string
	for rows.Next() {
		var id int
		err := rows.Scan(&id, &item)
		if err != nil {
			log.Panic(err)
		}
	}
	return item, nil
}
