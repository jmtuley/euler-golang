package problem_5

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestProblem_5(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Problem 5")
}

var _ = Describe("SmallestNumberDivisibleByAll", func() {
  It("works for 5", func() {
    // 5 * 4 * 3 == 60
    Expect(SmallestNumberDivisibleByAll(5)).To(Equal(60))
  })

  It("works for 10", func() {
    Expect(SmallestNumberDivisibleByAll(10)).To(Equal(2520))
  })

})
