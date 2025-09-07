package main

import "fmt"


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	list := [5]int{1, 2, 3, 4, 5}
	go func ()  {
		defer close(ch1)
		for _, value := range list {
			ch1 <- value
		}
	}()
	go func ()  {
		defer close(ch2)
		for value := range ch1 {
			ch2 <- value*2
		}
	}()
	for value := range ch2 {
		fmt.Println(value)
	}
}