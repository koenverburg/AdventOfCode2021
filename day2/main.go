package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

func forward(horizontal int, units int) int {
  return units + horizontal 
}

func up(depth int, units int) int {
  return depth - units
}

func down(depth int, units int) int {
  return units + depth
}

func calculateResult(depth int, horizontal int) int {
  return horizontal * depth
}

func instructionParser(instruction string) (string, int) {
  s := strings.Split(instruction, " ")
  action := s[0]
  value, err := strconv.ParseInt(s[1], 10, 64)
  utils.CheckIfErr(err)

  return action, int(value)
}

func runner(depth int, horizontal int, rawInstruction string) (int, int) {
  action, value := instructionParser(rawInstruction)
	switch action {
	case "forward":
    horizontal = forward(horizontal, value)
  case "down":
    depth = down(depth, value)
  case "up":
    depth = up(depth, value)
	}

  return depth, horizontal
}

func main() {
  depth := 0
  horizontal := 0
  input := utils.ReadFromTextFile("day2/input.txt")
  // input := []string{
  //   "forward 5",
  //   "down 5",
  //   "forward 8",
  //   "up 3",
  //   "down 8",
  //   "forward 2",
  // }

  for i := 0; i < len(input); i++ {
    value1, value2 := runner(depth, horizontal, input[i])
    depth = value1
    horizontal = value2
  }

  fmt.Println(calculateResult(depth, horizontal))
}
