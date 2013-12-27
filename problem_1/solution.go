package problem_1

import (
	"fmt"
)

func main() {
	fmt.Println("Total is ", sum_of_threes_and_fives_below(1000))
}

func sum_of_threes_and_fives_below(max int) int {

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
	number_of_threes := (max - 1) / 3
	number_of_fives := (max - 1) / 5
	number_of_fifteens := (max - 1) / 15
	sum_of_threes := 3 * (number_of_threes * (number_of_threes + 1)) / 2
	sum_of_fives := 5 * (number_of_fives * (number_of_fives + 1)) / 2
	sum_of_fifteens := 15 * (number_of_fifteens * (number_of_fifteens + 1)) / 2
	sum = sum_of_threes + sum_of_fives - sum_of_fifteens

	// return sum_of_fives + sum_of_threes
	return sum

}
