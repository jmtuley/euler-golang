package problem_4

import (
  "strconv"
)

func Solution() int {
  return LargestPalindromicProduct(999)
}

func LargestPalindromicProduct(factorMax int) int {
  max := 0
  for i := factorMax; i > 0; i-- {
    if max > i*i {
      return max
    }

    for j := i; j > 0; j-- {
      if isPalindrome(i*j) && i*j > max {
        max = i * j
      }
    }
  }
  return 0
}

func isPalindrome(value int) bool {
  s := strconv.Itoa(value)
  length := len(s)

  for i := 0; i <= length/2; i++ {
    if s[i] != s[length-i-1] {
      return false
    }
  }

  return true
}
