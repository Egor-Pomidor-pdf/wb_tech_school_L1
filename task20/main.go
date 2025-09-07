package main

import (
	"fmt"
)

func reverse(runes []rune, start int, end int) {
  for i, j := start, end; i < j; i, j = i + 1, j - 1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
}

func reverseText(text string) string {
  textRunes := []rune(text)
  reverse(textRunes, 0, len(textRunes) - 1)

  start := 0
  for i := 0; i <= len(textRunes); i++ {
    if i == len(textRunes) || textRunes[i] == ' ' {
      reverse(textRunes, start, i - 1)
      start = i + 1
    }
  }
  return string(textRunes)

}
 


func main() {
  fmt.Println(reverseText("egor popisal v shtan"))
}


// func reverseText(text string) string {
//   words := strings.Split(text, " ")
//   for i, j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1{
//     words[i], words[j] = words[j], words[i]
//   }
//   return strings.Join(words, " ")
// }



// Разработать программу, которая переворачивает порядок слов в строке.

// Пример: входная строка:

// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».