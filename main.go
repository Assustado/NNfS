package main

import (
	"math/rand"
	"time"
)

func main() {
	// m1 := newMatrix(2, 3)
	// m1.data[0].data[0] = 1
	// m1.data[0].data[1] = 2
	// m1.data[0].data[2] = 3
	// m1.data[1].data[0] = 4
	// m1.data[1].data[1] = 5
	// m1.data[1].data[2] = 6

	// m2 := newMatrix(3, 2)
	// m2.data[0].data[0] = 7
	// m2.data[0].data[1] = 8
	// m2.data[1].data[0] = 9
	// m2.data[1].data[1] = 10
	// m2.data[2].data[0] = 11
	// m2.data[2].data[1] = 12

	// m3 := multiplyMatrix(m1, m2)

	// m3.print()

	now := time.Now()
	rand.Seed(now.Unix())
	nn := initNN([]int{2, 1, 2})
	nn.print()

	inputs := newMatrix(2, 4)
	targets := newMatrix(2, 4)

	inputs.data[0].data[0] = 1
	inputs.data[1].data[0] = 1
	targets.data[0].data[0] = 1
	targets.data[1].data[0] = 0

	inputs.data[0].data[1] = 0
	inputs.data[1].data[1] = 0
	targets.data[0].data[1] = 1
	targets.data[1].data[1] = 0

	inputs.data[0].data[2] = 0
	inputs.data[1].data[2] = 1
	targets.data[0].data[2] = 0
	targets.data[1].data[2] = 1

	inputs.data[0].data[3] = 1
	inputs.data[1].data[3] = 0
	targets.data[0].data[3] = 1
	targets.data[1].data[3] = 1

	for i := 0; i < 1; i++ {
		nn.train(inputs, targets)
	}
}
