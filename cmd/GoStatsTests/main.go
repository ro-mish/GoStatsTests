package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	// Anscombe's Quartet data sets
	dataSets := []struct {
		x []float64
		y []float64
	}{
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		},
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		},
		{
			x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		},
		{
			x: []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
			y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
		},
	}

	for i, dataSet := range dataSets {
		// Calculate means
		xMean, _ := stats.Mean(dataSet.x)
		yMean, _ := stats.Mean(dataSet.y)

		// Calculate slope using covariance and variance
		var numerator, denominator float64
		for j := range dataSet.x {
			numerator += (dataSet.x[j] - xMean) * (dataSet.y[j] - yMean)
			denominator += (dataSet.x[j] - xMean) * (dataSet.x[j] - xMean)
		}
		slope := numerator / denominator

		// Calculate intercept using y = mx + b
		intercept := yMean - slope*xMean

		fmt.Printf("Dataset %d: Slope = %.4f, Intercept = %.4f\n", i+1, slope, intercept)
	}
}
