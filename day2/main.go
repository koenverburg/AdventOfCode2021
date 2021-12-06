package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

func down(aim int, units int) int {
  return aim + units
}

func up(aim int, units int) int {
  return aim - units
}

func forward(horizontal int, depth int, aim int, units int) (int, int) {
  newHorizontal := horizontal + units
  depth += units * aim

  return newHorizontal, depth
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

func runner(aim int, depth int, horizontal int, rawInstruction string) (int, int, int) {
  action, value := instructionParser(rawInstruction)
	switch action {
  case "up":
    aim = up(aim, value)
  case "down":
    aim = down(aim, value)
	case "forward":
    horizontal, depth = forward(horizontal, depth, aim, value)
	}

  return depth, horizontal, aim
}

func main() {
  aim := 0
  depth := 0
  horizontal := 0
  // input := utils.ReadFromTextFile("day2/input.txt")
  input := []string{
    "forward 5",
    "down 5",
    "forward 8",
    "up 3",
    "down 8",
    "forward 2",
  }

  for i := 0; i < len(input); i++ {
    value1, value2, value3 := runner(aim, depth, horizontal, input[i])
    depth = value1
    horizontal = value2
    aim = value3
  }

  fmt.Println(calculateResult(depth, horizontal))
}
