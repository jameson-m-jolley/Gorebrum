package net

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
type Node struct {
	inputs     *mat.VecDense
	weights    *mat.VecDense
	bias       float64
	activation string
}

// method for computing the output of a node
func (n Node) Compute_node() float64 {
	// hash of actavation functions sored in the ./activation.go file this is done to sapport diffrent actavation functions throught out the network of layer level
	return Node_Activation_Functions[n.activation](mat.Dot(n.inputs, n.weights)+n.bias, n)
}

// makes a new node
func Newnode(inputs, weights []float64, bias float64, activation string) *Node {
	return &Node{
		inputs:     mat.NewVecDense(len(inputs), inputs),
		weights:    mat.NewVecDense(len(weights), weights),
		bias:       bias,
		activation: activation}
}

func (n Node) Display_info() {
	fmt.Printf(`
---------------------------------------------------------------------

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

`, n.inputs, n.weights, n.bias, n.Compute_node(), n.activation)

}
