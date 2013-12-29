package problem_1

func Solution() int {
  return sumOfThreesAndFivesBelow(1000)
}

func sumOfThreesAndFivesBelow(max int) int {

  sum := 0

  // // NAIVE METHOD:
  // // loop over all integers from 1 to max, adding to sum if divisible by 3 or 5.
  // for i := 1; i < max; i++ {
  //  if (i%3 == 0) || (i%5 == 0) {
  //    sum += i
  //  }
  // }

  // // SLIGHTLY LESS NAIVE METHOD
  // // Loop over all multiples of 3 or 5, avoiding double-counting
  // for i := 3; i < max; i += 3 {
  // 	sum += i
  // }

  // for j := 5; j < max; j += 5 {
  // 	if j%3 != 0 {
  // 		sum += j
  // 	}
  // }

  // // FAST METHOD
  // // Use Little Gauß's formula to sum multiples of three and five; then subtract off multiples of 15.
  numberOfThrees := (max - 1) / 3
  numberOfFives := (max - 1) / 5
  numberOfFifteens := (max - 1) / 15
  sumOfThrees := 3 * (numberOfThrees * (numberOfThrees + 1)) / 2
  sumOfFives := 5 * (numberOfFives * (numberOfFives + 1)) / 2
  sumOfFifteens := 15 * (numberOfFifteens * (numberOfFifteens + 1)) / 2
  sum = sumOfThrees + sumOfFives - sumOfFifteens

  return sum
}
