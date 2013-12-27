package problem_1

import "testing"

func TestSumOfFivesAndThrees(t *testing.T) {
	sum := sum_of_threes_and_fives_below(3)
	if sum != 0 {
		t.Errorf("Sum was %v; wanted %v", sum, 0)
	}

	sum = sum_of_threes_and_fives_below(5)
	if sum != 3 {
		t.Errorf("Sum was %v; wanted %v", sum, 3)
	}

	sum = sum_of_threes_and_fives_below(6)
	if sum != 8 {
		t.Errorf("Sum was %v; wanted %v", sum, 8)
	}

	sum = sum_of_threes_and_fives_below(15)
	if sum != 45 {
		t.Errorf("Sum was %v; wanted %v", sum, 45)
	}

	// 3 + 6 + 9 + 5 == 23
	sum = sum_of_threes_and_fives_below(10)
	if sum != 23 {
		t.Errorf("Sum was %v; wanted %v", sum, 23)
	}

	sum = sum_of_threes_and_fives_below(1000)
	if sum != 233168 {
		t.Errorf("Sum was %v; wanted %v", sum, 233168)
	}

}
