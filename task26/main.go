package main

import (
	"fmt"
	"strings"
)

func checkUnic(s string) bool {
	s = strings.ToLower(s)
	flag := true
	m := make(map[rune]bool)
	for _, r := range s {
		_, ok := m[r]
		if ok {
			flag = false
			break
		}
		m[r] = true
	}
	if !flag {
		return false
	}
	return true
}

func main() {
	fmt.Println(checkUnic("saA"))
}