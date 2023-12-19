package services

import "strings"

func Count(s, w string) int {
	if w == "" {
		return 0 
	}

	words := strings.Fields(s) 

	count := 0
	for _, word := range words {
		if word == w {
			count++
		}
	}

	return count
}

func main4() {
	s := "apple orange banana apple mango apple"
	w := "apple"
	result := Count(s, w)
	println("Count of", w, "in", s, ":", result)
}
