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
	defer DbConn.Close()
	// prepareDB(DbConn)
	return DbConn
}

// func prepareDB(DbConn *sql.DB) {
// 	insertDataSQLStmt := `SELECT * FROM "public"."menu" LIMIT 100`
// 	rows, err := DbConn.Query(insertDataSQLStmt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id int
// 		var item string
// 		err := rows.Scan(&id, &item)
// 		if err != nil {
// 			log.Panic(err)
// 		}
// 		log.Println(id, item)
// 	}
// }
