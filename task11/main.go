package main

import "fmt"


func main() {
	mapInt := make(map[int]bool)
	slice1 := []int{1,2,3}
	slice2 := []int{2,3,4}
	var resultSlice []int

for _, val := range slice1 {
		mapInt[val] = true
}
for _, val := range slice2 {
	if mapInt[val] {
		resultSlice = append(resultSlice, val)
	}
}



fmt.Println(resultSlice)



	


}