package main

import (
	"math"
	"gonum.org/v1/gonum/mat"
	"time"
	"math/rand"
	"fmt"
	"log"
)

func main() {
	// Define our input attributes.
	input := mat.NewDense(3, 4, []float64{
		1.0, 0.0, 1.0, 0.0,
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
	})

	// Define our labels.
	labels := mat.NewDense(3, 1, []float64{1.0, 1.0, 0.0})

	// Define our network architecture and learning parameters.
	config := NeuralNetConfig{
		architecture: []int{4, 1, 1},
		trainingIterations: 5000,
		learningRate: 0.3,
		goodEnoughError: 0.005,
	}

	// Train the neural network.
	network := newNetwork(config)
	if err := network.Train(input, labels); err != nil {
		log.Fatal(err)
	}

	network.Print()
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}

type NeuralNetConfig struct {
	architecture []int
	goodEnoughError float64
	learningRate float64
	trainingIterations int
}

type NeuralNet struct {
	// user provided
	config NeuralNetConfig
	inputs *mat.Dense
	// trained
	weights *mat.Dense
	biases *mat.Dense
	//activations *mat.Dense
	// magic
	outputs *mat.Dense
}

func newNetwork(config NeuralNetConfig) *NeuralNet {
	return &NeuralNet{config: config}
}

func (nn *NeuralNet) init() *NeuralNet {
	rand.Seed(time.Now().UnixNano())

	if len(nn.config.architecture) < 3 {
		return nil
	}

	var rawWeights [][]float64
	var rawBiases [][]float64
	numberOfLayers := len(nn.config.architecture)
	rawWeights = make([][]float64, numberOfLayers)
	for i, layerSize := range nn.config.architecture {
		if i == 0 {
			// First layer (input) has not weights. Adding empty weights entry keeps clean index logic between for weight and architecture layers
			continue
		}
		previousLayerSize := nn.config.architecture[i-1]

		// Create arrays to hold trainable parameters
		rawWeights[i] = make([]float64, previousLayerSize*layerSize) // each neuron on this layer has previousLayerSize weighs
		rawBiases[i] = make([]float64, layerSize) // one bias per neuron

		// Initialize weights and biases at random
		fillRandom(rawWeights[i])
		fillRandom(rawBiases[i])
		nn.weights = mat.NewDense(layerSize, previousLayerSize, rawWeights[i])
		nn.biases = mat.NewDense(layerSize, previousLayerSize, rawBiases[i])
	}

	return nn
}

func fillRandom(a []float64) {
	for i := range a {
		a[i] = rand.Float64()
	}
}

func (nn *NeuralNet) Train(inputs *mat.Dense, expectedOutputs *mat.Dense) error {
	nn.init()
	nn.outputs = mat.NewDense(0,0, nil)

	for i := 0; i < nn.config.trainingIterations; i++ {
		//nn.feedForward()
		//deltas := nn.backProp()
		//nn.adjust(deltas)
	}
	return nil
}

func (nn *NeuralNet) Print() {
	// Print network weights
	f := mat.Formatted(nn.weights, mat.Prefix(" "))
	fmt.Printf("\nweights = %v\n\n", f)

	// Print network biases
	f = mat.Formatted(nn.biases, mat.Prefix(" "))
	fmt.Printf("\nbiases = %v\n\n", f)
}