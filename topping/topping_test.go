package topping

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestFetchMenu(t *testing.T) {
	_, err := fetchMenu()
	if err != nil {
		t.Error(err)
	}
}
func TestGetToppings(t *testing.T) {
	_, err := getToppings()
	if err != nil {
		t.Error(err)
	}
}

func TestGetToppingByName(t *testing.T) {
	toppingType := [...]string{"None", "Powdered Sugar"}
	for _, str := range toppingType {
		result, err := getToppingByName(str)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%+v\n", result)
	}
}
