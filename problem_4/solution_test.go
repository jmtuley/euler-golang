package problem_4

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestProblem_4(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Problem 4")
}

var _ = Describe("LargestPalindromicProduct", func() {
  It("works for numbers at most 11", func() {
    Expect(LargestPalindromicProduct(11)).To(Equal(121))
  })

  It("works for single-digit numbers", func() {
    Expect(LargestPalindromicProduct(9)).To(Equal(9))
  })

  It("works for the given sample", func() {
    Expect(LargestPalindromicProduct(99)).To(Equal(9009))
  })
})
