package main

type neuralNetwork struct {
	layers    []layer
	output    []float64
	numLayers int
}

type layer struct {
	neurons    []neuron
	numNeurons int
	output     []float64
}

type neuron struct {
	weights []float64
	bias    float64
	output  float64
}

func initNN(numLayers int, layersSize []int) neuralNetwork {
	// initialize neural network
	nn := neuralNetwork{
		numLayers: len(layersSize),
		layers:    []layer{},
	}

	il := layer{
		output: make([]float64, layersSize[0]),
	}
	nn.layers = append(nn.layers, il)
	//initialize hidden and output layers
	for i := 1; i < len(layersSize); i++ {
		l := layer{
			neurons:    []neuron{},
			numNeurons: layersSize[i],
			output:     make([]float64, layersSize[i]),
		}
		for j := 0; j < layersSize[i]; j++ {
			l.neurons = append(l.neurons, neuron{
				weights: randomWeights(layersSize[i-1]),
				bias:    0.0,
				output:  0.0,
			})
		}
		nn.layers = append(nn.layers, l)
	}

	return nn
}

func (nn *neuralNetwork) run(input []float64) {
	nn.layers[0].output = input
	for i := 1; i < nn.numLayers; i++ {
		nn.layers[i].run(nn.layers[i-1].output)
	}
	nn.output = nn.layers[nn.numLayers-1].output
}

func (layer *layer) run(input []float64) {
	for i, neuron := range layer.neurons {
		neuron.output = sigmoid(dot(input, neuron.weights) + neuron.bias)
		layer.output[i] = neuron.output
	}
}
