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

func runPart2(list []int) int {
  sumList := []int{}
  startingPoint := 0

  for startingPoint < len(list) {
    idx := startingPoint
    if (idx+1 >= len(list) || idx+2 >= len(list)) {
      break
    }

    value1 := list[idx]
    value2 := list[idx+1]
    value3 := list[idx+2]

    sum := value1 + value2 + value3
    sumList = append(sumList, sum)

    startingPoint++
  }

  // fmt.Println(sumList)

  return run(sumList)
}


func main() {
  // input := utils.ReadFromTextFileInt("day1/input.txt")
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
  }

  result := runPart2(input)
  fmt.Printf("The result of the first day is '%v'\n", result)
}
