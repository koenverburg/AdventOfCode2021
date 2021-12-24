const {runner} = require('../part1')

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

describe('day3', () => {
  beforeAll(() => {
    console.log = () => {}
  })

  test('part1', () => {
    const result = runner(input)
    expect(result.gamma).toBe(22)
    expect(result.eps).toBe(9)
    expect(result.consumption).toBe(198)
  })
})
