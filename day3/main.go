package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

func convertToInt(binary string) uint64 {
  value, err := strconv.ParseUint(binary, 2, 32)
  utils.CheckIfErr(err)
  return value
}


func sumBool(b ...bool) int {
  n := 0
  for _, v := range b {
    if v {
      n++
    }
  }
  return n
}

func flippingBits(b ...bool) string {
  r := []string{}
  for _, v := range b {
    if v {
      r = append(r, "1")
    } else {
      r = append(r, "0")
    }
  }

  return strings.Join(r, "")
}

func createGamma(lines []string) []bool {
  boolSlice := [][]bool{}
  for _, line := range lines {
    bucket := []bool{}
    for _, char := range strings.Split(line, "") {
      if char == "1" {
        bucket = append(bucket, true)
      } else {
        bucket = append(bucket, false)
      }
    }
    boolSlice = append(boolSlice, bucket)
  }
  
  result := []bool{}
  for _, x := range boolSlice {
    result = append(result, sumBool(x...) > len(x) / 2)
  }

  return result
}

func flipToEpsilon(lines []bool) []bool {
  r := []bool{}
  for _, line := range lines {
    r = append(r, !line)
  }
  return r
}

func runner(lines []string) {
  gammaList := createGamma(lines)
  epsilonList := flipToEpsilon(gammaList)

  fmt.Println(gammaList, flippingBits(gammaList...))
  fmt.Println(epsilonList, flippingBits(epsilonList...))

  gamma := convertToInt(flippingBits(gammaList...))
  epsilon := convertToInt(flippingBits(epsilonList...))

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
