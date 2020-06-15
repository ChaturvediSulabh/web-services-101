package topping

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func getToppings(fileName string) ([]Topping, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var s Sample
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	var toppings []Topping
	for i := 0; i < len(s); i++ {
		toppings = append(toppings, s[i].Topping)
	}
	return toppings, nil
}

func getToppingByName(fileName, toppingType string) ([]ToppingResponse, error) {
	var s Sample
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
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
