package main

import (
	"fmt"
	"sync"
)

type PesonAge struct {
	name string
	age uint
}

var mu sync.Mutex
var wg sync.WaitGroup

var sl = []PesonAge{{"DEDA", 18}, {"BABA", 52}, {"DEDA", 18}, {"OLEG", 180}}

func main() {
	mapTheBest := map[string]int{}

	for _, val := range sl {
		wg.Add(1)
		go func (p PesonAge)  {
			defer wg.Done()
			mu.Lock()
			mapTheBest[p.name] = int(p.age)
			mu.Unlock()
		}(val)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены")
	fmt.Println(mapTheBest)

}