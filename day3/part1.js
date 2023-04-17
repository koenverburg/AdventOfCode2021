const fs = require('fs')
const path = require('path')

// const input = [
//   "00100",
//   "11110",
//   "10110",
//   "10111",
//   "10101", 
//   "01111",
//   "00111",
//   "11100",
//   "10000",
//   "11001",
//   "00010",
//   "01010"
// ]

const buffer = fs.readFileSync(path.join(__dirname, './input.txt'))
const input = buffer.toString().split('\n')

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

  const gammaBinary = createBinaryDigit(pivot)
  const epsBinary = createBinaryDigitFlipped(pivot)

  console.log(gammaBinary)
  console.log(epsBinary)

  const gamma = parseInt(gammaBinary.join(''), 2) 
  const eps = parseInt(epsBinary.join(''), 2) 

  console.log(`result : ${gamma} * ${eps} = ${gamma * eps}`)

  return {
    eps,
    gamma,
    consumption: gamma * eps
  }
}

runner(input)
module.exports = { runner }
