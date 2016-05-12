// Comment here will create Pakcage Overview
package math

import (
	"fmt"
)

// Fucntion from custom Math package
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	fmt.Println("custom Average func")
	return total / float64(len(xs))
}

func ExampleAverage() {
	fmt.Println(Average([]float64{1.3, 2.5, 10.0}))
	// Output: olleh
}
