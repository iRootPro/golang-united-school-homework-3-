package homework

func average(input [15]float32) (result float32) {

	for i := 0; i < 15; i++ {
		result += input[i]
	}

	return result / 15
}
