package main

import (
	"testing"
	"time"
)

func TestMainRegression(t *testing.T) {
	start := time.Now()

	// Call the main function or the specific function that performs the regression
	main()

	duration := time.Since(start)
	t.Logf("Regression process took %s", duration)
}
