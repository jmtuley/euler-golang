package problem_3

func Solution() uint64 {
  return LargestPrimeFactorOf(600851475143)
}

func LargestPrimeFactorOf(target uint64) uint64 {
  generator := primeGenerator()

  var max uint64
  for current := <-generator; current <= target; current = <-generator {
    if target%current == 0 {
      max = current
      target /= max
    }
  }

  return max
}

func primeGenerator() chan uint64 {
  output := make(chan uint64)
  primes := make([]uint64, 0)
  current := uint64(2)

  go func() {
    for {
      for ; !testPrimality(current, primes); current++ {
      }

      primes = append(primes, current)

      output <- current
      current++
    }
  }()

  return output
}

func testPrimality(target uint64, primes []uint64) bool {
  for _, p := range primes {
    if target%p == 0 {
      return false
    }
  }

  return true
}
