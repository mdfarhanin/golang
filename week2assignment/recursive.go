package main

import "encoding/hex"

var i = 10

func Recursive(a string) string {
	src := []byte(a)
	if a == "" {
		return "no string"
	}
	for i < 1 {
		return hex.EncodeToString(src)
	}
	i--
	return hex.EncodeToString(src) + Recursive(a)
}
