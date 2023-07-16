package net

import (
	"gonum.org/v1/gonum/mat"
)

type Gorebrum struct {
	input          *mat.VecDense
	layers         []Layer
	output         *mat.VecDense
	x              int
	y              int
	out_dimensions int
}
