package topping

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"web-services-101/cors"
	"web-services-101/database"
)

//SetupRoutes ...
func SetupRoutes(apiBasePath string) {
	toppings, err := getToppings()
	if err != nil {
		log.Fatal(err)
	}
	toppingByNameHandler := http.HandlerFunc(toppingHandler())
	toppingListHandler := http.HandlerFunc(toppingsHandler(toppings))
	http.Handle(fmt.Sprintf("%s%s", apiBasePath, "/topping/"), cors.Middleware(toppingByNameHandler))
	http.Handle(fmt.Sprintf("%s%s", apiBasePath, "/toppings"), cors.Middleware(toppingListHandler))
}

func toppingHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPathSegments := strings.Split(r.URL.Path, "/")
		toppingType := urlPathSegments[len(urlPathSegments)-1]
		myTopping, err := getToppingByName(toppingType)
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
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			insertSQLStmt := `INSERT INTO "public"."menu" (item) VALUES ($1)`
			_, err = database.SetupDB().ExecContext(ctx, insertSQLStmt, string(data))
			if err != nil {
				log.Panic(err)
			}
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
