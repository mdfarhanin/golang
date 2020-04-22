package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Operation struct {
	Operation string
	Units     Units
}

type Units struct {
	One int
	Two int
}

func CalculateFromFile() {
	data, _ := ioutil.ReadFile("supportFiles//calculateinput.json")
	var opr []Operation
	json.Unmarshal(data, &opr)
	add := 0
	subtract := 0
	multiply := 0
	divide := 0
	for _, v := range opr {
		switch v.Operation {
		case "Add":
			add = v.Units.One + v.Units.Two
		case "Subtract":
			subtract = v.Units.One - v.Units.Two
		case "Multiply":
			multiply = v.Units.One * v.Units.Two
		case "Divide":
			divide = v.Units.One / v.Units.Two
		}
	}

	d1 := []byte("Add: " + strconv.Itoa(add) + "\n" + "Subtract: " + strconv.Itoa(subtract) + "\n" + "Mutiply: " + strconv.Itoa(multiply) + "\n" + "Divide: " + strconv.Itoa(divide))
	err := ioutil.WriteFile("calculateoutput.txt", d1, 0644)
	if err != nil {
		panic(err)
	}

}
