package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

func convertToInt(binary string) int {
  value, err := strconv.Atoi(binary) // , 2, 64)
  utils.CheckIfErr(err)
  // fmt.Println("i8", strconv.FormatInt(int64(value), 2))

  return value
}


func createBoolList(lines []string) []string {
  list := []string {}
  for _, line := range lines {
    fmt.Println(line, strings.Count(line, "1"), strings.Count(line, "0"))
    if(strings.Count(line, "1") > len(line) / 2) {
      list = append(list, "1")
    } else {
      list = append(list, "0")
    }
  }
  return list
}

func flipToEpsilon(lines []string) []string {
  list := []string{}
  for _, line := range lines {
    if(line == "1") {
      list = append(list, "0")
    } else {
      list = append(list, "1")
    }
  }
  return list
}


func runner(lines []string) {
  gammaList := createBoolList(lines)
  epsilonList := flipToEpsilon(gammaList)

  // fmt.Println(gammaList)
  // fmt.Println(epsilonList)
  fmt.Println(strings.Join(gammaList, "")[7:])
  fmt.Println(strings.Join(epsilonList, "")[7:])

  gamma := convertToInt(strings.Join(gammaList[7:], ""))
  epsilon := convertToInt(strings.Join(epsilonList[7:], ""))

  fmt.Printf("The result of gamma rate is %v and epsilon rate is %v, which brings the power consumption to %v\n", gamma, epsilon, gamma * epsilon)
}

func main() {
  input := []string {
    "00100",
    "11110",
    "10110",
    "10111",
    "10101", 
    "01111",
    "00111",
    "11100",
    "10000",
    "11001",
    "00010",
    "01010",
  }
  runner(input)
}
