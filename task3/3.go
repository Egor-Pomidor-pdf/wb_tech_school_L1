package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)



func main() {
	numbers := make(chan int)

	if len(os.Args) < 2 {
		fmt.Println("напиши скок workers должно быть")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Wrrong numWorkers")
		return
	}

	for i := 0; i < numWorkers; i++ {
		go worker(i, numbers)
	}

	count := 0
	for {
		numbers <- count
		count++
		time.Sleep(100 * time.Millisecond)
	}
		
}


func worker(id int, numbers chan int) {
	for num := range numbers {
		fmt.Printf("Worker id=%d work with %d\n", id, num)
	}
}
