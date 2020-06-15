package main

import (
	"log"
	"net/http"
	"web-services-101/topping"
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
	const apibasePath = "/api"
	topping.SetupRoutes(apibasePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
