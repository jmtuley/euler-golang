package problem_3

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestProblem_3(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Problem 3")
}

var _ = Describe("primeGenerator", func() {
  It("gives the first few prime numbers", func() {
    c := primeGenerator()

    Expect(<-c).To(Equal(uint64(2)))
    Expect(<-c).To(Equal(uint64(3)))
    Expect(<-c).To(Equal(uint64(5)))
    Expect(<-c).To(Equal(uint64(7)))
    Expect(<-c).To(Equal(uint64(11)))
    Expect(<-c).To(Equal(uint64(13)))
  })
})

var _ = Describe("LargestPrimeFactorOf", func() {
  It("finds the largest prime factor of 2", func() {
    Expect(LargestPrimeFactorOf(2)).To(Equal(uint64(2)))
  })

  It("finds the largest prime factor of 6", func() {
    Expect(LargestPrimeFactorOf(6)).To(Equal(uint64(3)))
  })

  It("finds the largest prime factor of 2310", func() {
    Expect(LargestPrimeFactorOf(2310)).To(Equal(uint64(11)))
  })
})
