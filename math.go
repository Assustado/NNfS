package main

import (
	"math"
)

const e float64 = 2.7182818284590

func sigmoid(input float64) float64 {
	return (1 / (1 + (math.Pow(e, input*-1.0))))
}

func derivativeSigmoid(input float64) float64 {
	output := (input) * (1 - (input))
	return output
}

func cost(a1, a2 []float64) []float64 {
	output := make([]float64, len(a1))
	for i := range a1 {
		output[i] += (a1[i] - a2[i]) * (a1[i] - a2[i])
	}
	return output
}
