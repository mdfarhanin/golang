package main

import (
	"fmt"
	"strconv"
	"time"
)

type dealer struct {
	dealername string
	cars       []car
}

type owner struct {
	firstname string
	lastname  string
	address   string
	dob       string
	ownercar  car
}

type car struct {
	caryear   int
	carmake   string
	carmodel  string
	imei      string
	carstatus string
	insurance string
}

func (t *car) Init(year int, make string, model string, insurance string) {
	t.caryear = year
	t.carmake = make
	t.carmodel = model
	t.insurance = insurance
}

func (t *car) setStatus(carstatus1 string) {
	t.carstatus = carstatus1
}

func Cars() {
	mapcars := make(map[string]car)

	//Declaring and Initializing stock
	newcars := make([]car, 5)
	for i, _ := range newcars {
		newcars[i].Init(2020, "Toyota", "Camry SE", "GEICO")
		newcars[i].carstatus = "NEW"
		newcars[i].imei = "Toyota" + strconv.Itoa(i)
		mapcars[newcars[i].imei] = newcars[i]
	}

	usedcars := make([]car, 5)
	for i, _ := range usedcars {
		usedcars[i].Init(2020, "Mitsubishi", "Model X", "Progressive")
		usedcars[i].carstatus = "USED"
		usedcars[i].imei = "Mitsubishi" + strconv.Itoa(i)
		mapcars[usedcars[i].imei] = usedcars[i]
	}

	// Initializing owners
	o1 := owner{"David", "Coppper", "123 street", "04/10/1980", mapcars["Toyota3"]}
	o2 := owner{"Kleet", "Field", "asd street", "05/10/1982", mapcars["Mitsubishi1"]}
	o3 := owner{"Pi", "Null", "qwe street", "05/10/1984", mapcars["Toyota4"]}

	d1 := dealer{dealername: "CarMax", cars: newcars}
	d2 := dealer{dealername: "CarVana", cars: usedcars}

	fmt.Println("**********Dealer 1*************")
	time.Sleep(2 * time.Second)
	fmt.Println("Dealer 1: ", d1)
	fmt.Println("**********Dealer 2*************")
	time.Sleep(2 * time.Second)
	fmt.Println("Dealer 2: ", d2)

	fmt.Println("**********Owner 1*************")
	time.Sleep(2 * time.Second)
	fmt.Println("Owner 1: ", o1)
	fmt.Println("**********Owner 2*************")
	time.Sleep(2 * time.Second)
	fmt.Println("Owner 2: ", o2)
	fmt.Println("**********Owner 3*************")
	time.Sleep(2 * time.Second)
	fmt.Println("Owner 3: ", o3)

}
