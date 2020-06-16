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
	prepareDB(DbConn)
	return DbConn
}

func prepareDB(DbConn *sql.DB) {
	// createTableSQLStmt := `
	// CREATE TABLE menu (
	// 	id SERIAL PRIMARY KEY,
	// 	item JSON NOT NULL
	// )
	// `
	insertDataSQLStmt := `
	INSERT INTO "public"."menu"
	(item)
	values
	('{
			"id": "0001",
			"type": "donut",
			"name": "Cake",
			"ppu": 0.55,
			"topping": [
				{
					"id": "5001",
					"type": "None"
				},
				{
					"id": "5002",
					"type": "Glazed"
				},
				{
					"id": "5005",
					"type": "Sugar"
				},
				{
					"id": "5007",
					"type": "Powdered Sugar"
				},
				{
					"id": "5006",
					"type": "Chocolate with Sprinkles"
				},
				{
					"id": "5003",
					"type": "Chocolate"
				},
				{
					"id": "5004",
					"type": "Maple"
				}
			]
		}'),
		('{
			"id": "0002",
			"type": "donut",
			"name": "Raised",
			"ppu": 0.55,
			"topping": [
				{
					"id": "5001",
					"type": "None"
				},
				{
					"id": "5002",
					"type": "Glazed"
				},
				{
					"id": "5005",
					"type": "Sugar"
				},
				{
					"id": "5003",
					"type": "Chocolate"
				},
				{
					"id": "5004",
					"type": "Maple"
				}
			]
		}'),
		('{
			"id": "0003",
			"type": "donut",
			"name": "Old Fashioned",
			"ppu": 0.55,
			"topping": [
				{
					"id": "5001",
					"type": "None"
				},
				{
					"id": "5002",
					"type": "Glazed"
				},
				{
					"id": "5003",
					"type": "Chocolate"
				},
				{
					"id": "5004",
					"type": "Maple"
				}
			]
		}')'
	`
	res, err := DbConn.Exec(insertDataSQLStmt)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
