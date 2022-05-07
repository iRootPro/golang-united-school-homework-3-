package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	result_arr := []string{}

	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		append(result_arr, input[k])
	}
	return result_arr
}
