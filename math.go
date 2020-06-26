package main

import (
	"math"
	"math/rand"
)

const e float64 = 2.7182818284590

func dot(a1, a2 []float64) float64 {
	output := 0.0
	for i := range a1 {
		output += a1[i] * a2[i]
	}
	return output
}

func randomWeights(size int) []float64 {
	output := make([]float64, size)
	for i := range output {
		output[i] = (rand.Float64() * 2) - 1
	}
	return output
}

func sigmoid(input float64) float64 {
	return (1 / (1 + (math.Pow(e, input*-1))))
}
