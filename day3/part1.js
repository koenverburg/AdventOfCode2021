// const fs = require('fs')

const input = [
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
  "01010"
]

const inputWidth = input[0].length;
const pivot = new Array(input[0].length).fill(0)

const getColumn = (input, columnIndex) => [...input][columnIndex]
const countValue = input => input.filter(char => char === '1').length
const createBinaryDigit = list => list.map(b => b > inputWidth ? 1 : 0)
const createBinaryDigitFlipped = list => list.map(b => b > inputWidth ? 0 : 1)

const walkColumn = (input, colIndex) => {
  const column = []

  for (let index = 0; index < input.length; index++) {
    column[index] = getColumn(input[index], colIndex)
  }
  
  return countValue(column)
}

const runner = (input) => {
  for (let index = 0; index < inputWidth; index++) {
    pivot[index] = walkColumn(input, index)
  }

  const gamma = createBinaryDigit(pivot)
  const eps = createBinaryDigitFlipped(pivot)

  console.log({r: gamma, result: parseInt(gamma.join(''), 2) })
  console.log({r: eps, result: parseInt(eps.join(''), 2) })

  return {
    eps: parseInt(eps.join(''), 2),
    gamma: parseInt(gamma.join(''), 2)
  }
}

runner(input)
module.exports = { runner }
