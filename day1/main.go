package main

import (
	"fmt"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

func run(list []int) int {
  increasesCount := 0

  for i := 0; i < len(list); i++ {
    if i == 0 {
      continue
    }

    index := i - 1 

    previousValue := list[index]

    if (list[i] > previousValue) {
      // fmt.Println(list[i])
      increasesCount = increasesCount + 1
    }
  }

  return increasesCount
}

func main() {
  input := utils.ReadFromTextFile("day1/input.txt")
	// input := []int{
	// 	199,
	// 	200,
	// 	208,
	// 	210,
	// 	200,
	// 	207,
	// 	240,
	// 	269,
	// 	260,
	// 	263,
  // }

  result := run(input)
  fmt.Printf("The result of the first day is '%v'\n", result)
}
