package main

import "fmt"

func main() {
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapTemp := make(map[int][]float64)
	for _, val := range list {
		group := int(val) / 10 * 10
		mapTemp[group] = append(mapTemp[group], val)
	}

	fmt.Println(mapTemp)

}