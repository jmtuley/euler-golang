package problem_1

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestProblem_1(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "sumOfThreesAndFivesBelow")
}

var _ = Describe("sumOfThreesAndFivesBelow", func() {
  It("works for 3", func() {
    sum := sumOfThreesAndFivesBelow(3)
    Expect(sum).To(Equal(0))
  })

  It("works for 5", func() {
    sum := sumOfThreesAndFivesBelow(5)
    Expect(sum).To(Equal(3))
  })

  It("works for 6", func() {
    sum := sumOfThreesAndFivesBelow(6)
    Expect(sum).To(Equal(8))
  })

  It("works for 10", func() {
    sum := sumOfThreesAndFivesBelow(10)
    Expect(sum).To(Equal(23))
  })

  It("works for 15", func() {
    sum := sumOfThreesAndFivesBelow(15)
    Expect(sum).To(Equal(45))
  })

  It("works for 1000", func() {
    sum := sumOfThreesAndFivesBelow(1000)
    Expect(sum).To(Equal(233168))
  })
})
