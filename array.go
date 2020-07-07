package main

import (
	"fmt"
	"math/rand"
)

type array struct {
	len  int
	data []float64
}

func newArray(len int) array {
	output := array{
		len:  len,
		data: make([]float64, len),
	}
	return output
}

func randomArray(len int, min, max float64) array {
	output := array{
		len:  len,
		data: make([]float64, len),
	}
	for i := range output.data {
		output.data[i] = (rand.Float64() * (max - min)) + min
	}
	return output
}

func multiplyArray(a1, a2 array) float64 {
	output := 0.0
	for i := range a1.data {
		output += a1.data[i] * a2.data[i]
	}
	return output
}

func (a1 *array) print() {
	fmt.Print("[")
	for i := 0; i < a1.len; i++ {
		fmt.Print(a1.data[i])
		if i != a1.len-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
