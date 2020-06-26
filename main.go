package main

import (
	"fmt"
)

func main() {
	nn := initNN(2, []int{2, 2})
	fmt.Println(nn)
	nn.run([]float64{1, 1})
	fmt.Println(nn)
}
