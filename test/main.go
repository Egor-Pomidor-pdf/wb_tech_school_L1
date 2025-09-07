package main

import "fmt"

func findMin(sliceFloat []float64) (int, float64) {
	minValue := sliceFloat[0]
	minValueIndex := 0

	for ind, value := range sliceFloat {
		if value < minValue {
			minValue = value
			minValueIndex = ind
		}
	}
	return  minValueIndex, minValue
}

func removeFromeSlice(sliceFloat []float64, pos int)[]float64 {
	return append(sliceFloat[:pos], sliceFloat[pos+1:]...)
}


func main() {
	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapTemp := map[int][]float64{}
		for range list {
			i, minVal := findMin(list)
			tenGroup := int(minVal) / 10 % 10
			for _, val := range list {
				if ((int(val) / 10 % 10) == tenGroup) {
					mapTemp[tenGroup*10] = append(mapTemp[tenGroup*10], val)
				}
			}
			list = removeFromeSlice(list, i)

	
		}
		
		
		fmt.Println(mapTemp)

}
