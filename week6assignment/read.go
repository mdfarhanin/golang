package main

import (
	"fmt"
	"io/ioutil"
)

func Reading() {
	data, err := ioutil.ReadFile("supportFiles//example.json")
	check(err)
	fmt.Println("*****Reading JSON and Printing: ", string(data))

	data, err = ioutil.ReadFile("supportFiles//example.csv")
	check(err)
	fmt.Println("*****Reading CSV and Printing: ", string(data))

	data, err = ioutil.ReadFile("supportFiles//example.txt")
	check(err)
	fmt.Println("*****Reading Text and Printing: ", string(data))

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
