package main

import (
	"flag"
	"log"
	"net/http"
	"web-services-101/database"
	"web-services-101/topping"

	_ "github.com/lib/pq"
)

//Sample ...
type Sample []struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Topping Topping `json:"topping"`
}

//Topping ...
type Topping []struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//ToppingResponse ...
type ToppingResponse struct {
	Name string
	Cost float64
}

func main() {
	port := flag.String("PORT", "5000", "PORT to be used to run the application")
	flag.Parse()
	addr := ":" + *port
	database.SetupDB()
	const apibasePath = "/api"
	topping.SetupRoutes(apibasePath)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
