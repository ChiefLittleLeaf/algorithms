package algorithms

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum

	// Sum() this version takes the first element in the slice[0] and as long as the slice continues to not be empty will add the remaining numbers to running total (can cause stack overflow) *Recursion*
	//func Sum(numbers []int ) int{
	//	if len(numbers) == 0 {
	//		return 0
	//	}
	//	return numbers[0] + Sum(numbers[1:])
}
