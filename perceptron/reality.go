package main

import "math/rand"

func importData(size int) []LabeledDatum {
	var points = make([]LabeledDatum, size)
	for i := 0; i < size; i++ {
		points[i].Datum = createPoint()
		points[i].Label = isAboveLine(points[i].Datum, realFunction)
	}
	return points
}

func createPoint() Datum {
	return Datum{
		Features: []float64{
			rand.Float64()*100,
			rand.Float64()*100,
		},
	}
}

func realFunction(x float64) float64 {
	return 1.5*x + 1.7
}

func isAboveLine(point Datum, f func(float64) float64) Label {
	x := point.Features[0]
	y := point.Features[1]
	if y > f(x) {
		return positiveClass
	}
	return negativeClass
}