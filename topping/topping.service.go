package topping

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"web-services-101/cors"
)

//SetupRoutes ...
func SetupRoutes(apiBasePath string) {
	var fileName *string = flag.String("dataFile", "data/sample.json", "JSON file with sample data")
	flag.Parse()
	toppings, err := getToppings(*fileName)
	if err != nil {
		return
	}
	toppingByNameHandler := http.HandlerFunc(toppingHandler(*fileName))
	toppingListHandler := http.HandlerFunc(toppingsHandler(toppings))
	http.Handle(fmt.Sprintf("%s%s", apiBasePath, "/topping/"), cors.Middleware(toppingByNameHandler))
	http.Handle(fmt.Sprintf("%s%s", apiBasePath, "/toppings"), cors.Middleware(toppingListHandler))
}

func toppingHandler(data string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("HERE")
		urlPathSegments := strings.Split(r.URL.Path, "/")
		toppingType := urlPathSegments[len(urlPathSegments)-1]
		log.Println(urlPathSegments, toppingType)
		myTopping, err := getToppingByName(data, toppingType)
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
