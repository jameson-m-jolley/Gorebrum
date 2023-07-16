package net

import (
	"math"
)

/*
▄▀█ █▀▀ ▀█▀ █ █░█ ▄▀█ ▀█▀ █ █▀█ █▄░█   █▀▀ █░█ █▄░█ █▀▀ ▀█▀ █ █▀█ █▄░█ █▀
█▀█ █▄▄ ░█░ █ ▀▄▀ █▀█ ░█░ █ █▄█ █░▀█   █▀░ █▄█ █░▀█ █▄▄ ░█░ █ █▄█ █░▀█ ▄█

we are using a dict to store the funtions. each function will be stored under a key that is a string so that each node can have a key to a actavation function
all
*/

// rectified linear: takes a float64 as var x and returns 0 if input < 0 else returns var x
func ReLU(x float64, n Node) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

/*
█▀▀ ▀▄▀ █▀█ █▀█ █▄░█ █▀▀ █▄░█ ▀█▀ █ ▄▀█ ▀█▀ █ █▀█ █▄░█
██▄ █░█ █▀▀ █▄█ █░▀█ ██▄ █░▀█ ░█░ █ █▀█ ░█░ █ █▄█ █░▀█
*/

// raw exponentiation via math.E
func ExpE(x float64, n Node) float64 {
	return math.Pow(math.E, x)
}

// raw exponentiation via math.SqrtE
func ExpSqrtE(x float64, n Node) float64 {
	return math.Pow(math.SqrtE, x)
}

// raw exponentiation via PI
func ExpPi(x float64, n Node) float64 {
	return math.Pow(math.Pi, x)
}

// raw exponentiation via PI
func ExpSqrtPi(x float64, n Node) float64 {
	return math.Pow(math.SqrtPi, x)
}

/*
█░█ ▄▀█ █▀ █░█ █▀▄▀█ ▄▀█ █▀█
█▀█ █▀█ ▄█ █▀█ █░▀░█ █▀█ █▀▀
*/

// dict to map funtions for the Node obj, this is done for dinamic calls to diffrent func.
// all func must be plan func and not methods with the parma(x float64, n Node) or cannot be placed inside the map
var Node_Activation_Functions map[string]func(float64, Node) float64 = map[string]func(float64, Node) float64{
	"ReLU":      ReLU,
	"ExpE":      ExpE,
	"ExpSqrtE":  ExpSqrtE,
	"ExpPi":     ExpPi,
	"ExpSqrtPi": ExpSqrtPi,
}

// Layer_based
// soft max
func SoftMax(x float64, l layer) float64 {
	return x
}

/*
█░█ ▄▀█ █▀ █░█ █▀▄▀█ ▄▀█ █▀█
█▀█ █▀█ ▄█ █▀█ █░▀░█ █▀█ █▀▀
*/

var layer_Activation_Functions map[string]func(float64, layer) float64 = map[string]func(float64, layer) float64{
	"SoftMax": SoftMax,
}
