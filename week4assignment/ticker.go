package main

import (
	"fmt"
	"time"
)

func Ticker() {

	ticker := time.NewTicker(1 * time.Second)
	i := 1
	for t := range ticker.C {
		fmt.Println("Tick at ", t)
		i = i + 1
		if i == 12 {
			break
		}
	}
	fmt.Println("Ticker stopped")
}
