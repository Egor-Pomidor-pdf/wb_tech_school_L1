package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	ch := make(chan int)
	if len(os.Args) < 2 {
		fmt.Println("напиши скок time должно быть")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error of N(time a work of programm)")
		return
	}
	stopFromTimer := time.After(time.Duration(N) * time.Second)
	counter := 0

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	for {
		select {
		case <-stopFromTimer:
			close(ch)
			return
		default:
			ch <- counter
			counter++
			time.Sleep(100 * time.Millisecond)
		}
	}

}
