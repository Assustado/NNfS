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
	numNeurons int
	input      matrix
	output     matrix
}

func initNN(layersSize []int) neuralNetwork {
	// initialize neural network
	nn := neuralNetwork{
		numLayers:    len(layersSize),
		layers:       []layer{},
		learningRate: 0.1,
	}

	il := layer{
		output:     newMatrix(1, layersSize[0]),
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

func (nn *neuralNetwork) train(input, target matrix) {

	nn.layers[0].output = input
	for i := 1; i < nn.numLayers; i++ {
		nn.layers[i].run(nn.layers[i-1].output)
	}
	nn.output = nn.layers[nn.numLayers-1].output

	// output layer

	nn.layers[2].newWeights = newMatrix(nn.layers[2].weights.rows, nn.layers[2].weights.cols)

	// fmt.Println("================================")
	// fmt.Println(nn.layers[2].input.data[0].data[0])
	// fmt.Println("================================")

	auxD_lastlayer_n0 := (target.data[0].data[0] - nn.output.data[0].data[0]) * (derivativeSigmoid(nn.layers[2].output.data[0].data[0]))
	grad := auxD_lastlayer_n0 * nn.layers[1].output.data[0].data[0]

	nn.layers[2].newWeights.data[0].data[0] = nn.layers[2].weights.data[0].data[0] + grad

	// hidden layer

	nn.layers[1].newWeights = newMatrix(nn.layers[1].weights.rows, nn.layers[1].weights.cols)

	auxD_hiddenlayer_n0 := (auxD_lastlayer_n0) * (nn.layers[2].weights.data[0].data[0]) * (derivativeSigmoid(nn.layers[1].output.data[0].data[0]))
	grad = auxD_hiddenlayer_n0 * nn.layers[0].output.data[0].data[0]

	nn.layers[1].newWeights.data[0].data[0] = nn.layers[1].weights.data[0].data[0] + grad

	auxD_hiddenlayer_n1 := (auxD_lastlayer_n0) * (nn.layers[2].weights.data[0].data[0]) * (derivativeSigmoid(nn.layers[1].output.data[0].data[0]))
	grad = auxD_hiddenlayer_n1 * nn.layers[0].output.data[1].data[0]

	nn.layers[1].newWeights.data[0].data[1] = nn.layers[1].weights.data[0].data[1] + grad

	nn.layers[2].weights = copyMatrix(nn.layers[2].newWeights)

	nn.layers[1].weights = copyMatrix(nn.layers[1].newWeights)

	fmt.Println(nn.output.data[0].data[0] - target.data[0].data[0])
}

func (nn *neuralNetwork) run(input matrix) {
	nn.layers[0].output = copyMatrix(input)
	for i := 1; i < nn.numLayers; i++ {
		nn.layers[i].run(nn.layers[i-1].output)
	}
	nn.output = nn.layers[nn.numLayers-1].output
}

func (layer *layer) run(input matrix) {
	layer.output = multiplyMatrix(layer.weights, input)
	layer.input = copyMatrix(layer.output)

	for i := 0; i < len(layer.output.data); i++ {
		for j := 0; j < len(layer.output.data[i].data); j++ {
			layer.output.data[i].data[j] = sigmoid(layer.output.data[i].data[j])
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
