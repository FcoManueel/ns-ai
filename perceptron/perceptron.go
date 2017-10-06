package perceptron


type Label float64

type ActivationFunction func(float64) float64

type Perceptron struct {
	Weights []float64
	Activate ActivationFunction
}

func NewPerceptron (inputSize int) Perceptron {
	if inputSize < 1 {
		inputSize = 1
	}
	weights := make([]float64, inputSize+1) // +1 to account for bias weight
	return Perceptron{
		Weights: weights,
		Activate: sign,
	}
}

func sign(n float64) float64 {
	if n < 0 {
		return -1
	}
	return 1
}

type Datum struct {
	Features []float64
}

type LabeledDatum struct {
	Datum
	Label Label
}

func (p *Perceptron) Train(trainingData []LabeledDatum) {
	for true {
		hasReachedOptimum := true // Up to this point this is an assumption, which will be determined after the following loop
		//fmt.Println("====== Starting new iteration ======")
		for _, datum := range trainingData {
			predictedLabel := p.Predict(datum.Datum)
			isCorrectPrediction := datum.Label == predictedLabel
			if !isCorrectPrediction {
				hasReachedOptimum = false
				//fmt.Printf("   EFORE: w0: %.7f \t w1: %.7f \t b: %.7f\n", p.Weights[0], p.Weights[1], p.Weights[2])
				p.Weights = p.adjustWeights(datum)
				//fmt.Printf("   AFTER: w0: %.7f \t w1: %.7f \t b: %.7f\n\n\n", p.Weights[0], p.Weights[1], p.Weights[2])
			}
		}
		if hasReachedOptimum { break }
	}
}

func (p *Perceptron) adjustWeights(datum LabeledDatum) []float64 {
	for i := 0; i < len(p.Weights) - 1; i++ {
		p.Weights[i] += datum.Datum.Features[i] * float64(datum.Label)
		//fmt.Printf("-------> Delta[%d] = %+w \n", i, datum.Datum.Features[i] * float64(datum.Label) *0.1)
	}
	p.Weights[len(p.Weights) - 1] += 1 * float64(datum.Label)
	return p.Weights
}


func (p *Perceptron) Predict(data Datum) Label {
	var sum float64
	for i, feature := range data.Features {
		sum += feature * p.Weights[i]
	}
	// Weights has an extra element than Features. It's the bias
	sum += p.Weights[len(data.Features)]

	return Label(p.Activate(sum))
}