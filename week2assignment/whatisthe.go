package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int{
	return 42
}

func init() {
	WhatIsThe = 0
}

func Main1(){
	if WhatIsThe == 0 {
		fmt.Println("Its all a lie")
	}
}