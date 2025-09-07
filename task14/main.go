package main

import "fmt"

func typeVariable(i interface{}) {
	switch i.(type){
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("This is bool")
	case chan int:
        fmt.Println("This is chan int")
    case chan string:
        fmt.Println("This is chan string")
    case chan bool:
        fmt.Println("This is chan bool")
	default:
		fmt.Printf("This is unknown type: %T\n", i)
	}

}


func main() {
	ch := make(chan string)
	typeVariable(ch)
}