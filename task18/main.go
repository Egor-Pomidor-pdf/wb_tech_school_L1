package main

import (
	"fmt"
	"sync"
)

type Counter struct {
  count int
}

func (c * Counter) increment() {
  c.count += 1
}

func (c * Counter) showCount() {
  fmt.Println(c.count)
}

func main() {
  var wg sync.WaitGroup
  var mu sync.Mutex
  var c Counter
  for i:= 0; i < 3; i++ {
    wg.Add(1)
    go func() {
      for i:= 0; i < 5; i++ {
        mu.Lock()
        c.increment()
        mu.Unlock()
    }
    wg.Done()
  }()
}
wg.Wait()
c.showCount()
}