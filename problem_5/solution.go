package problem_5

type PrimeFactorization struct {
  PrimePowers []int
}

func Solution() int {
  return SmallestNumberDivisibleByAll(20)
}

func SmallestNumberDivisibleByAll(factorMax int) int {
  factorizations := make([]PrimeFactorization, factorMax)

  for i := 1; i <= factorMax; i++ {
    factorizations[i-1] = newPrimeFactorization(i)
  }

  maxLength := maxFactorizationLength(factorizations)

  powers := make([]int, 0)
  for i := 0; i < maxLength; i++ {
    max := 0
    for j := 0; j < factorMax; j++ {
      if factorizations[j].powerForIndexedPrime(i) > max {
        max = factorizations[j].powerForIndexedPrime(i)
      }
    }

    powers = append(powers, max)
  }

  generator := primeGenerator()

  product := 1
  for i := 0; i < len(powers); i++ {
    p := <-generator

    for j := 0; j < powers[i]; j++ {
      product *= p
    }
  }
  return product
}

func maxFactorizationLength(factorizations []PrimeFactorization) int {
  m := 0
  for _, factorization := range factorizations {
    if len(factorization.PrimePowers) > m {
      m = len(factorization.PrimePowers)
    }
  }

  return m
}

func (pf PrimeFactorization) powerForIndexedPrime(index int) int {
  if index >= len(pf.PrimePowers) {
    return 0
  } else {
    return pf.PrimePowers[index]
  }
}

func newPrimeFactorization(value int) PrimeFactorization {
  return PrimeFactorization{
    PrimePowers: primeFactorPowers(value),
  }
}

func primeFactorPowers(value int) []int {
  generator := primeGenerator()

  powers := make([]int, 0)

  for p := <-generator; p <= value; p = <-generator {
    power := 0
    for v := value; v%p == 0; power++ {
      v /= p
    }

    powers = append(powers, power)
  }

  return powers
}

func primeGenerator() chan int {
  output := make(chan int)
  primes := make([]int, 0)
  current := int(2)

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

func testPrimality(target int, primes []int) bool {
  for _, p := range primes {
    if target%p == 0 {
      return false
    }
  }

  return true
}
