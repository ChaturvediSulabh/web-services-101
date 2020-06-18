package topping

import (
	"encoding/json"
	"errors"
	"web-services-101/database"
)

func fetchMenu() ([]Sample, error) {
	rows := database.GetAllData(database.SetupDB())
	var s []Sample
	for _, data := range rows {
		var d Sample
		err := json.Unmarshal([]byte(data), &d)
		if err != nil {
			return nil, err
		}
		s = append(s, d)
	}
	return s, nil
}

func getToppings() ([]Topping, error) {
	s, err := fetchMenu()
	if err != nil {
		return nil, err
	}
	var toppings []Topping
	for i := 0; i < len(s); i++ {
		toppings = append(toppings, s[i].Topping)
	}
	return toppings, nil
}

func getToppingByName(toppingType string) ([]ToppingResponse, error) {
	s, err := fetchMenu()
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
