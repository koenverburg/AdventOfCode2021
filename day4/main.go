package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/koenverburg/AdventOfCode2021/utils"
)

type Coordinate struct {
  X int
  Y int
  Value int
}

type Board struct {
  matrix [][]int
  markedNumbers []Coordinate
}

func convertToIntSlice(input []string) []int {
  result := []int{}
  for i := 0; i < len(input); i++ {
    number, err := strconv.ParseInt(input[i], 10, 64)
    utils.CheckIfErr(err)

    result = append(result, int(number))
  }
  return result
}

func createMatrix(input string) [][]int {
	rows := strings.Split(input, " ")
	board := [][]int{}

	offset := 0
	for i := 0; i < 5; i++ {
		row := []int{}
		numbers := rows[offset:offset+5]

		for j := 0; j < 5; j++ {
			value, _ := strconv.ParseInt(numbers[j], 10, 64)
			row = append(row, int(value))
		}

		offset += 5
		board = append(board, row)
	}

  return board
}

func IncludesNumber(board Board, number int) (Coordinate, bool) {
  for i := 0;i < len(board.matrix); i++ {
    for j := 0;j < len(board.matrix[i]); j++ {
      boardNumber := board.matrix[i][j]
      if boardNumber == number {
        coordinate := Coordinate{
          Value: number,
          X: i,
          Y: j,
        }
        return coordinate, true
      } else {
        continue
      }
    }
  }
  return Coordinate{}, false
}

func CheckForWin(board Board) Board {
  increasesCount := 0
  count := []Coordinate{}
  list := board.markedNumbers
  for i := 0; i < len(list); i++ {
    if i == 0 {
      continue
    }

    fmt.Println(list[i])

    index := i - 1 

    previousValue := list[index].X

    if (list[i].X > previousValue) {
      fmt.Println(list[i])
      increasesCount = increasesCount + 1
      count = append(count, list[i])
    }

    if increasesCount == 5 {
      break
    }
  }

  fmt.Println(count)
  
  return board
}

func main() {
  drawNumbers := []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}
  input := "22 13 17 11 0 8 2 23 4 24 21 9 14 16 7 6 10 3 18 5 1 12 20 15 19"
  matrix := createMatrix(input)

  var board Board
  board.matrix = matrix

  for i := 0; i < len(drawNumbers); i++ {
    location, ok := IncludesNumber(board, drawNumbers[i])
    if (ok) {
      board.markedNumbers = append(board.markedNumbers, location)
    }
  }

  winner := CheckForWin(board)
  fmt.Println(winner.matrix)

  // fmt.Println(board.markedNumbers)
}
