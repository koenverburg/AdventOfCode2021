package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CheckIfErr(err error) {
 if err != nil {
      panic(err)
  }
}

func ReadFromTextFile(path string) []int {
	file, err := os.Open(path)
  CheckIfErr(err)

  scanner := bufio.NewScanner(file)
  var items []int
  for scanner.Scan() {
    value := scanner.Text()
    linevalue, err := strconv.ParseInt(value, 10, 64)
    CheckIfErr(err)
    items = append(items, int(linevalue))
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return items
}

func ReadFromTextFileInt(path string) []int {
	file, err := os.Open(path)
  CheckIfErr(err)

  scanner := bufio.NewScanner(file)
  var items []int
  for scanner.Scan() {
    value := scanner.Text()
    linevalue, err := strconv.ParseInt(value, 10, 64)
    CheckIfErr(err)
    items = append(items, int(linevalue))
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return items
}
