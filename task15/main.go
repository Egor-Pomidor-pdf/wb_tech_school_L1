package main

import (
	"fmt"
)


func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
  v := createHugeString(1 << 10)
  result := make([]byte, 100)

  copy(result, v[:100])
  fmt.Println(string(result))
}

func main() {
  someFunc()
}

// В изначальном коде была проблема утечки памяти, так как juststring ссылалась на часть огормной строки, которая не могла быть очищена сборщиком мусора из-за того, что на нее ссылались