package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type Sample []struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Batters struct {
		Batter []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"batter"`
	} `json:"batters"`
	Topping []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"topping"`
}

func main() {
	dataFile := flag.String("data-file", "data/sample.json", "JSON file with sample data")
	data, err := ioutil.ReadFile(*dataFile)
	if err != nil {
		log.Fatal(err)
	}
	var s Sample
	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n", s)
}