package net

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/*
▄▀█ █▀▀ ▀█▀ █ █░█ ▄▀█ ▀█▀ █ █▀█ █▄░█   █▀▀ █░█ █▄░█ █▀▀ ▀█▀ █ █▀█ █▄░█ █▀
█▀█ █▄▄ ░█░ █ ▀▄▀ █▀█ ░█░ █ █▄█ █░▀█   █▀░ █▄█ █░▀█ █▄▄ ░█░ █ █▄█ █░▀█ ▄█

we are using a dict to store the funtions. each function will be stored under a key that is a string so that each node can have a key to a actavation function
all
*/
// dict to map funtions for the Node obj, this is done for dinamic calls to diffrent func.
// all func must be plan func and not methods with the parma(x float64, n Node) or cannot be placed inside the map
var Node_Activation_Functions map[string]func(float64, *Node) float64 = map[string]func(float64, *Node) float64{
	"NodeLU":    NodeLU,
	"ReLU":      ReLU,
	"Log":       Log,
	"ExpE":      ExpE,
	"ExpSqrtE":  ExpSqrtE,
	"ExpPi":     ExpPi,
	"ExpSqrtPi": ExpSqrtPi,
	"Collatz":   Collatz,
}

// rectified linear: takes a float64 as var x and returns 0 if input < 0 else returns var x
func ReLU(x float64, n *Node) float64 {
	if x < 0 {
		n.Output = 0
		return 0
	} else {
		n.Output = x
		return x
	}
}

// natral log: takes a float64 as var x and returns 0 if input < 0 else returns var log10(x)
func Log(x float64, n *Node) float64 {
	n.Output = math.Log10(x)
	return n.Output
}

// raw exponentiation via math.E
func ExpE(x float64, n *Node) float64 {
	n.Output = math.Pow(math.E, x)
	return n.Output
}

func NodeLU(x float64, n *Node) float64 {
	n.Output = x
	return n.Output
}

//experamental activation

// raw exponentiation via math.SqrtE
func ExpSqrtE(x float64, n *Node) float64 {
	n.Output = math.Pow(math.SqrtE, x)
	return n.Output
}

// raw exponentiation via PI
func ExpPi(x float64, n *Node) float64 {
	n.Output = math.Pow(math.Pi, x)
	return n.Output
}

// raw exponentiation via PI
func ExpSqrtPi(x float64, n *Node) float64 {
	n.Output = math.Pow(math.SqrtPi, x)
	return n.Output
}

// the hailstone algorithom or Collatz conjecture https://en.wikipedia.org/wiki/Collatz_conjecture
func Collatz(x float64, n *Node) float64 {

	// we half to tuncate the val so that the val can be an int then at the end of the function the floating point desapales are added to the val
	newx := math.Round(x)
	dif := x - newx

	if math.Mod(newx, 2) == 0 {
		newx = newx / 2
	} else {
		newx = 3*newx + 1
	}

	return newx + dif

}

var Layer_Activation_Functions map[string]func(*mat.VecDense, Layer) = map[string]func(x *mat.VecDense, L Layer){
	"NA":      NA,
	"SoftMax": SoftMax,
}

// Layer_based
// soft max

// this function is intededed to be usesd for the Output of the forward pass so
// that the NN can be traned with the prodiction of how corect the model is
func SoftMax(x *mat.VecDense, L Layer) {

	exponentiation(0, x)
	// clip 0 vals
	normalize(0, x, x.Norm(1))

	for i := 0; i < x.Len(); i++ {

		if x.AtVec(i) < 1e-7 {
			x.SetVec(i, 1e-7)
		} else if x.AtVec(i) > 1-1e-7 {
			x.SetVec(i, 1-1e-7)
		}
	}

}

// this is part of the soft max activation
func exponentiation(i int, x *mat.VecDense) {
	if i < x.Len() {
		x.SetVec(i, math.Pow(math.E, x.AtVec(i)))
		i++
		exponentiation(i, x)
	}
}

// this is part of the soft max activation
func normalize(i int, x *mat.VecDense, norm float64) {
	if i < x.Len() {
		x.SetVec(i, x.AtVec(i)/norm)
		i++
		normalize(i, x, norm)
	}
}

func NA(x *mat.VecDense, L Layer) {
}
