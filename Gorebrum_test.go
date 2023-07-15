package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

/*
███╗░░██╗░█████╗░██████╗░███████╗
████╗░██║██╔══██╗██╔══██╗██╔════╝
██╔██╗██║██║░░██║██║░░██║█████╗░░
██║╚████║██║░░██║██║░░██║██╔══╝░░
██║░╚███║╚█████╔╝██████╔╝███████╗
╚═╝░░╚══╝░╚════╝░╚═════╝░╚══════╝
*/
// this is the struct for creating nodes
// this will eventualy be a nural network
type node struct {
	inputs     *mat.VecDense
	weights    *mat.VecDense
	bias       float64
	activation string
}

/*
▄▀█ █▀▀ ▀█▀ █ █░█ ▄▀█ ▀█▀ █ █▀█ █▄░█   █▀▀ █░█ █▄░█ █▀▀ ▀█▀ █ █▀█ █▄░█ █▀
█▀█ █▄▄ ░█░ █ ▀▄▀ █▀█ ░█░ █ █▄█ █░▀█   █▀░ █▄█ █░▀█ █▄▄ ░█░ █ █▄█ █░▀█ ▄█

we are using a dict to store the funtions. each function will be stored under a key that is a string so that each node can have a key to a actavation function
*/
var activation_functions map[string]func(float64) float64 = map[string]func(float64) float64{
	"ReLU": ReLU,
}

// rectified linear: takes a float64 as var x and returns 0 if input < 0 else returns var x
func ReLU(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

// method for computing the output of a node
func (n node) compute_node() float64 {
	return activation_functions[n.activation](mat.Dot(n.inputs, n.weights) + n.bias)
}

// makes a new node
func newnode(inputs, weights []float64, bias float64, activation string) *node {
	L_inputs := mat.NewVecDense(len(inputs), inputs)
	L_weights := mat.NewVecDense(len(weights), weights)

	return &node{inputs: L_inputs, weights: L_weights, bias: bias, activation: activation}
}

func (n node) display_info() {
	fmt.Printf(`
	███╗░░██╗░█████╗░██████╗░███████╗
	████╗░██║██╔══██╗██╔══██╗██╔════╝
	██╔██╗██║██║░░██║██║░░██║█████╗░░
	██║╚████║██║░░██║██║░░██║██╔══╝░░
	██║░╚███║╚█████╔╝██████╔╝███████╗
	╚═╝░░╚══╝░╚════╝░╚═════╝░╚══════╝
	in: %v
	weights: %v
	bias: %f
	out: %f
	activation_function: %s
	`, n.inputs, n.weights, n.bias, n.compute_node(), n.activation)

}

/*
██╗░░░░░░█████╗░██╗░░░██╗███████╗██████╗░
██║░░░░░██╔══██╗╚██╗░██╔╝██╔════╝██╔══██╗
██║░░░░░███████║░╚████╔╝░█████╗░░██████╔╝
██║░░░░░██╔══██║░░╚██╔╝░░██╔══╝░░██╔══██╗
███████╗██║░░██║░░░██║░░░███████╗██║░░██║
╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝
*/

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/
func print_logo() {
	fmt.Print(`

░██████╗░░█████╗░██████╗░███████╗██████╗░██████╗░██╗░░░██╗███╗░░░███╗
██╔════╝░██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██║░░░██║████╗░████║
██║░░██╗░██║░░██║██████╔╝█████╗░░██████╦╝██████╔╝██║░░░██║██╔████╔██║
██║░░╚██╗██║░░██║██╔══██╗██╔══╝░░██╔══██╗██╔══██╗██║░░░██║██║╚██╔╝██║
╚██████╔╝╚█████╔╝██║░░██║███████╗██████╦╝██║░░██║╚██████╔╝██║░╚═╝░██║
░╚═════╝░░╚════╝░╚═╝░░╚═╝╚══════╝╚═════╝░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░╚═╝
Created by: Jameson Jolley
version: 1.0
`)
}

func main() {

	print_logo()
	anode := newnode([]float64{1.1, 2.3, 5.1}, []float64{1.8, 8.3, 7.1}, 4.5, "ReLU")
	fmt.Print(anode.bias)
	anode.display_info()
	fmt.Print(anode.compute_node())
}
