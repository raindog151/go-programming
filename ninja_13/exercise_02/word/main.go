package word

import "strings"

// UseCount returns a map of string->int representing the amount of times a word is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns an int representing non whitespace "words" in a string
func Count(s string) int {
	x := strings.Fields(s)
	return len(x)
}
