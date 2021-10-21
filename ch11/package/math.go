package math

func Average(inp ...float64) float64 {
	var sum float64
	var inp_size int
	for _, value := range inp {
		sum += value
	}
	return sum / float64(inp_size)
}
