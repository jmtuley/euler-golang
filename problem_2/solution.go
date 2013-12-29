package problem_2

func Solution() int {
  return EvenFibonacciSum(4e6)
}

func EvenFibonacciSum(max int) int {
  generator := fibGenerator()

  sum := 0
  current := 0
  for current <= max {
    next := <-generator
    if next <= max && next%2 == 0 {
      sum += next
    }

    current = next
  }

  return sum
}

func fibGenerator() chan int {
  back_two := 0
  back_one := 1

  output := make(chan int)

  go func() {
    for {
      current := back_two + back_one
      back_two = back_one
      back_one = current
      output <- current
    }
  }()

  return output
}
