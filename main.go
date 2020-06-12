package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

//Sample ...
type Sample []struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Batters Batters `json:"batters"`
	Topping Topping `json:"topping"`
}

//Batters ...
type Batters struct {
	Batter `json:"batter"`
}

//Batter ...
type Batter []struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//Topping ...
type Topping []struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func main() {
	dataFile := flag.String("dataFile", "data/sample.json", "JSON file with sample data")
	toppings := getToppings(*dataFile)
	http.HandleFunc("/topping", toppingHandler(toppings))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func toppingHandler(toppings []Topping) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			topping, err := json.Marshal(toppings)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(topping)
		}
	}
}

func getToppings(fileName string) []Topping {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var s Sample
	err = json.Unmarshal(data, &s)
	if err != nil {
		log.Panic(err)
	}
	var toppings []Topping
	for i := 0; i < len(s); i++ {
		toppings = append(toppings, s[i].Topping)
	}
	return toppings
}
