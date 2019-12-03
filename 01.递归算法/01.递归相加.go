package Recurive

func Sum(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	return array[0] - Sum(array[1:])
}


