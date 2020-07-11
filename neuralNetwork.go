package main

import "fmt"

type neuralNetwork struct {
	layers       []layer
	output       matrix
	numLayers    int
	learningRate float64
}

type layer struct {
	weights    matrix
	newWeights matrix
	bias       matrix
	newBias    matrix
	numNeurons int
	inputs     matrix
	outputs    matrix
}

func initNN(layersSize []int) neuralNetwork {
	// initialize neural network
	nn := neuralNetwork{
		numLayers:    len(layersSize),
		layers:       []layer{},
		learningRate: 0.1,
	}

	il := layer{
		numNeurons: layersSize[0],
		bias:       randomMatrix(layersSize[0], 1, 0.0, 1.0),
	}
	nn.layers = append(nn.layers, il)
	//initialize hidden and output layers
	for i := 1; i < len(layersSize); i++ {
		l := layer{
			weights:    randomMatrix(layersSize[i], layersSize[i-1], 0.0, 1.0),
			bias:       randomMatrix(layersSize[i], 1, 0.0, 1.0),
			numNeurons: layersSize[i],
		}
		nn.layers = append(nn.layers, l)
	}

	return nn
}

func (nn *neuralNetwork) train(inputs, targets matrix) {
	nn.run(inputs)
	nn.layers[nn.numLayers-1].outputs.print()

	// input of the last layer
	z := nn.layers[nn.numLayers-1].inputs
	// output of the last layer
	a := nn.layers[nn.numLayers-1].outputs

	err := subMatrix(a, targets)

	ds := newMatrix(z.rows, z.cols)

	for i, row := range z.data {
		for j, data := range row.data {
			ds.data[i].data[j] = derivativeSigmoid(data)
		}
	}

	fmt.Println("==================")
	err.print()
	fmt.Println("==================")
	ds.print()
	fmt.Println("==================")

	delta := multiplyMatrixElem(err, ds)

	nn.layers[nn.numLayers-1].newBias = delta

	a = nn.layers[nn.numLayers-2].outputs

	nn.layers[nn.numLayers-1].newWeights = multiplyMatrix(a.transpose(), delta)

	for i := nn.numLayers - 2; i > 0; i-- {
		z := nn.layers[i-1].inputs

		ds := newMatrix(z.rows, z.cols)

		for i, row := range z.data {
			for j, data := range row.data {
				ds.data[i].data[j] = derivativeSigmoid(data)
			}
		}

		wT := nn.layers[i-1].weights.transpose()

		wDelta := multiplyMatrix(delta, wT)

		delta = multiplyMatrixElem(wDelta, ds)

		nn.layers[i].newBias = delta

		a = nn.layers[i-1].outputs

		nn.layers[i-1].newWeights = multiplyMatrix(a.transpose(), delta)
	}

	N := inputs.rows

	for i := 1; i < nn.numLayers; i++ {

		alpha := 0.1 / float64(N)

		w := nn.layers[i].weights

		nw := nn.layers[i].newWeights

		// b := nn.layers[i].bias
		nbt := nn.layers[i].newBias.transpose()
		nb := newMatrix(nbt.rows, 1)

		for i, row := range nbt.data {
			for _, data := range row.data {
				nb.data[i].data[0] += data
			}
		}

		for i, row := range nw.data {
			for j, data := range row.data {
				nw.data[i].data[j] = data * alpha
			}
		}

		for i, row := range nb.data {
			for j, data := range row.data {
				nb.data[i].data[j] = data * alpha
			}
		}

		nn.layers[i].weights = subMatrix(w, nw)
	}
}

func (nn *neuralNetwork) run(inputs matrix) {

	nn.layers[0].outputs = copyMatrix(inputs)
	for i := 1; i < nn.numLayers; i++ {
		nn.layers[i].run(nn.layers[i-1].outputs)
	}
}

func (layer *layer) run(inputs matrix) {
	layer.outputs = multiplyMatrix(layer.weights, inputs)
	layer.inputs = copyMatrix(layer.outputs)

	for i := 0; i < len(layer.outputs.data); i++ {
		for j := 0; j < len(layer.outputs.data[i].data); j++ {
			layer.outputs.data[i].data[j] = sigmoid(layer.outputs.data[i].data[j])
		}
	}
}

func (nn *neuralNetwork) print() {
	fmt.Println()
	for x := 0; x < nn.numLayers; x++ {
		nn.layers[x].weights.print()
		fmt.Print("[")
		for y := 0; y < nn.layers[x].numNeurons; y++ {
			fmt.Print("O")
			if y != nn.layers[x].numNeurons-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
	fmt.Print("output: ")
	nn.output.print()
	fmt.Println()
}
