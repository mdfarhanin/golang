package main

func SubString(test string, start int, end int) []rune {
	var somechars []rune
	for i, char := range test {
		if i >= start && i < end+1 {
			somechars = append(somechars, char)
		}
	}
	return somechars
}
