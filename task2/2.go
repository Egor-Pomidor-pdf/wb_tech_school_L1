package main

import "fmt"


func main() {
	numbers := []int{2, 4, 6, 8, 10}
	res := make(chan int, len(numbers))

	for _, num := range numbers {
		go func(n int) {
			res <- n * n
		}(num)
	}

	for i :=0; i < len(numbers); i++ {
		fmt.Println(<-res)
	} 
		
}

