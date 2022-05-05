package mathematic

import (
	"fmt"
	"strconv"
)

func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	average, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total/float64(len(numbers))), 64)
	return average
}
