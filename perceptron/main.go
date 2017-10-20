package main

import (
	"fmt"
)

func main() {
	data := importData(1000)
	perceptron := NewPerceptron(2)

	fmt.Printf("Starting!\n Data: %+w", data)
	perceptron.Train(data)
	w0, w1, w2 := perceptron.Weights[0], perceptron.Weights[1], perceptron.Weights[2]
	fmt.Printf("Trained!\n w0: %.7f \t w1: %.7f \t bias: %.7f\n", w0, w1, w2)
	fmt.Printf("        \n slope: %.7f \t intersect: %.7f \n\n\n", -w0/w1, -w2/w1, w2)
}
