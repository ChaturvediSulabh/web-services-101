package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	dataFile := flag.String("dataFile", "data/sample.json", "JSON file with sample data")
	toppings := getToppings(*dataFile)
	http.HandleFunc("/topping", toppingsHandler(toppings))
	http.HandleFunc("/toppings/", toppingHandler(*dataFile))
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func toppingHandler(data string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathSegments := strings.Split(r.URL.Path, "/")
		toppingType := urlPathSegments[len(urlPathSegments)-1]
		myTopping, err := getToppingByName(toppingType, data)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		toppingsFound, err := json.Marshal(myTopping)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(toppingsFound)
		w.Header().Set("Content-Type", "application/json")
	}
}

func toppingsHandler(toppings []Topping) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var t Topping
			err = json.Unmarshal(data, &t)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Printf("Data received:\n%+v", t)
			w.WriteHeader(http.StatusCreated)
		default:
			topping, err := json.Marshal(toppings)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
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

func getToppingByName(toppingType string, fileName string) ([]ToppingResponse, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var s Sample
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	var r []ToppingResponse
	doesExists := false
	for _, st := range s {
		for _, t := range st.Topping {
			if toppingType == t.Type {
				if doesExists == false {
					doesExists = true
				}
				data := ToppingResponse{
					st.Name,
					st.Ppu,
				}
				r = append(r, data)
			}
		}
		if doesExists == false {
			return nil, errors.New("Topping Not Found")
		}
	}
	return r, nil
}
