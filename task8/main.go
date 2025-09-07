package main

import "fmt"

// "fmt"
// "os"
// "strconv"

func main() {
	bitSet(5, 1, 1)
	

	
}

func bitSet(num int64, pos int, digit int) {
	var resultNum int64
	fmt.Printf("Исходное число %d в битовом формате %064b \n", num, num)
	if (digit == 1) {
		resultNum = num | (1 << pos)
	} else {
		resultNum = num &^ (1 << pos)
	}
	fmt.Printf("Результирующее число %d в битовом формате %064b", resultNum, resultNum)

}



