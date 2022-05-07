package homework

func reverse(input []int64) (result []int64) {

	var copy_arr = []int64{}
	copy(input, copy_arr)

	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return
}
