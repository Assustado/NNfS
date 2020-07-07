package main

import (
	"fmt"
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
	nn := initNN([]int{2, 2, 2})
	nn.print()

	inputs := []matrix{}
	targets := []matrix{}

	input1 := newMatrix(2, 1)
	input1.data[0].data[0] = 1
	input1.data[1].data[0] = 1
	target1 := newMatrix(2, 1)
	target1.data[0].data[0] = 1
	target1.data[1].data[0] = 0

	input2 := newMatrix(2, 1)
	input2.data[0].data[0] = 0
	input2.data[1].data[0] = 0
	target2 := newMatrix(2, 1)
	target2.data[0].data[0] = 1
	target2.data[1].data[0] = 0

	input3 := newMatrix(2, 1)
	input3.data[0].data[0] = 0
	input3.data[1].data[0] = 1
	target3 := newMatrix(2, 1)
	target3.data[0].data[0] = 0
	target3.data[1].data[0] = 1

	input4 := newMatrix(2, 1)
	input4.data[0].data[0] = 1
	input4.data[1].data[0] = 0
	target4 := newMatrix(2, 1)
	target4.data[0].data[0] = 1

	inputs = append(inputs, input1)
	inputs = append(inputs, input2)
	inputs = append(inputs, input3)
	inputs = append(inputs, input4)

	targets = append(targets, target1)
	targets = append(targets, target2)
	targets = append(targets, target3)
	targets = append(targets, target4)

	for i := 0; i < 1000; i++ {
		trainData := rand.Int() % 1
		// fmt.Println(trainData)
		nn.train(inputs[trainData], targets[trainData])
		// nn.print()
	}
	nn.print()

	fmt.Println()
	nn.run(input1)
	nn.output.print()

	fmt.Println()
	nn.run(input2)
	nn.output.print()

	fmt.Println()
	nn.run(input3)
	nn.output.print()

	fmt.Println()
	nn.run(input4)
	nn.output.print()

}
