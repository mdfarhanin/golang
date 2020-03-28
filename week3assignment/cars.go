package main

import (
	"fmt"
	"math/rand"
)

type car struct {
	make  string
	model string
	year  int
	imei  int
	brand
	status
}

type brand struct {
	feature string
}

type status struct {
	carstatus string
}

func (t *car) Init(make1 string, model1 string, year1 int, feature1 string) {
	t.make = make1
	t.model = model1
	t.year = year1
	t.feature = feature1
	t.imei = rand.Intn(1000000000)
}

func (t *car) setStatus(carstatus1 string) {
	t.carstatus = carstatus1
}

func listCars(listmap map[int]car) {
	for _, value := range listmap {
		fmt.Println(value.imei, value.make, value.model, value.year, value.feature, value.status)
	}
}

func buySellCar(mapcars map[int]car, imei int, status string) map[int]car {
	car1 := mapcars[imei]
	car1.setStatus(status)
	mapcars[imei] = car1
	return mapcars
}

func Cars() {
	mapcars := make(map[int]car)

	//Declaring and Initializing 5 cars
	newcars := make([]car, 5)
	for i, _ := range newcars {
		newcars[i].Init("Toyota", "Camry SE", 2020, "Lane Tracing")
		newcars[i].setStatus("NEW")
		mapcars[newcars[i].imei] = newcars[i]
	}

	//Declaring and initializing car1
	car1 := new(car)
	car1.Init("Mercedez", "Benz", 2020, "Auto-Engine Off in Traffic")
	car1.setStatus("NEW")
	car1.imei = 2394823498
	mapcars[car1.imei] = *car1

	//Declaring and initializing car2
	car2 := new(car)
	car2.Init("Tesla", "Model 3", 2020, "Self-Drive")
	car2.setStatus("NEW")
	car2.imei = 3443634633
	mapcars[car2.imei] = *car2

	//Buying car2
	mapcars = buySellCar(mapcars, 3443634633, "BUY")

	//Selling car1
	mapcars = buySellCar(mapcars, 2394823498, "SELL")

	listCars(mapcars)
}
