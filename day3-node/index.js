// const fs = require('fs')

const mostCommonValue = '1' 
const leastCommonValue = '0'

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

const countValue = (input, valueToLookFor) => {
  const result = Object.values(input).reduce((count, current) => {
    if (current === valueToLookFor) {
      count++
    }
    return count
  }, 0)

  return result > input.length / 2
} 

const collect = (input, value) => {
  return input.map(i => countValue(i, value))
}

const createBinaryDigit = (input) => {
  return input.map(b => b ? '1' : '0').join('')
}

const runner = (input) => {
  // console.log('10110', countValue('10110', mostCommonValue))
  // console.log('01001', countValue('01001', leastCommonValue))

  const gammaCollection = collect(input, mostCommonValue)
  const binary = createBinaryDigit(gammaCollection)
  console.log(parseInt(binary, 2))

  // console.log(countValue(binary, mostCommonValue))
  // console.log(countValue(binary, leastCommonValue))

  // console.log(binary)
  // console.log(10110, parseInt('10110', 2))
  // console.log(parseInt(binary, 2))

  // const epsCollection = collect(input, leastCommonValue)
  // console.log(epsCollection)
}

runner(input)

// fs.readFileSync('./day3/input.txt')
