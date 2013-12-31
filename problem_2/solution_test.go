package problem_2

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestProblem_2(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Problem 2")
}

var _ = Describe("fibGenerator", func() {
  It("gives the first few Fibonacci numbers", func() {
    c := fibGenerator()

    Expect(<-c).To(Equal(1))
    Expect(<-c).To(Equal(2))
    Expect(<-c).To(Equal(3))
    Expect(<-c).To(Equal(5))
    Expect(<-c).To(Equal(8))
    Expect(<-c).To(Equal(13))
  })
})

var _ = Describe("EvenFibonacciSum", func() {
  It("works with a limit of 1", func() {
    Expect(EvenFibonacciSum(1)).To(Equal(0))
  })

  It("works with a limit of 2", func() {
    Expect(EvenFibonacciSum(2)).To(Equal(2))
  })

  It("works with a limit of 3", func() {
    Expect(EvenFibonacciSum(3)).To(Equal(2))
  })

  It("works with a limit of 8", func() {
    Expect(EvenFibonacciSum(8)).To(Equal(10))
  })
})
