package perceptron

import (
	"fmt"
)

func main() {
	data := importData(100000)
	perceptron := NewPerceptron(2)

	//fmt.Printf("Starting!\n Data: %+w", data)
	perceptron.Train(data)
	fmt.Printf("Trained!\n w0: %.7f \t w1: %.7f \t bias: %.7f\n\n\n", perceptron.Weights[0], perceptron.Weights[1], perceptron.Weights[2])
}
