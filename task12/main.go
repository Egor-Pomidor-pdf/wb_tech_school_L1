package main

import "fmt"


func main() {
	set := make(map[string]bool)
	var resultSlice []string
	sliceStr := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range sliceStr {
			if !set[v] {
			set[v] = true
			resultSlice = append(resultSlice, v)
			}
	}
	fmt.Println(resultSlice)
}