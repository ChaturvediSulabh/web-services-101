package topping

import (
	"encoding/json"
	"errors"
	"log"
	"web-services-101/database"
)

func getToppings() ([]Topping, error) {
	data, err := database.GetAllData(database.SetupDB())
	if err != nil {
		return nil, err
	}
	var s Sample
	err = json.Unmarshal([]byte(data), &s)
	if err != nil {
		return nil, err
	}
	var toppings []Topping
	for i := 0; i < len(s); i++ {
		toppings = append(toppings, s[i].Topping)
	}
	log.Println(toppings)
	return toppings, nil
}

func getToppingByName(toppingType string) ([]ToppingResponse, error) {
	data, err := database.GetAllData(database.SetupDB())
	if err != nil {
		return nil, err
	}
	var s Sample
	err = json.Unmarshal([]byte(data), &s)
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
