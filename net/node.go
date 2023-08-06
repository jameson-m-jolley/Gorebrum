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
	weights      *mat.VecDense
	bias         float64
	activation   string
	output       float64
	parent_layer *Layer
	index        int
}

// geters

func (n Node) Get_input() *mat.VecDense {
	if n.parent_layer.index == 0 {
		return n.parent_layer.parent_network.input
	} else {
		return n.parent_layer.parent_network.layers[n.parent_layer.index-1].output
	}
}

func (n Node) Get_weights() *mat.VecDense {
	return n.weights
}

func (n Node) Get_bias() float64 {
	return n.bias
}

func (n Node) Get_activation() string {
	return n.activation
}

func (n Node) Get_output() float64 {
	return n.output
}

func (n Node) Get_parent_layer() *Layer {
	return n.parent_layer
}

func (n *Node) Set_weights(weights *mat.VecDense) {
	n.weights = weights
}

func (n *Node) Set_bias(bias float64) {
	n.bias = bias
}

func (n *Node) Set_activation(activation string) {
	n.activation = activation
}

func (n *Node) Set_output(output float64) {
	n.output = output
	n.parent_layer.output.SetVec(n.index, n.output)
}

// method for computing the output of a node
func (n *Node) Compute_node() {
	//updates the inputs
	// hash of actavation functions sored in the ./activation.go file this is done to sapport diffrent actavation functions throught out the network of layer level
	n.Set_output(Node_Activation_Functions[n.activation](mat.Dot(n.Get_input(), n.weights)+n.bias, n))
}

// makes a new node
func (L *Layer) New_Node(weights []float64, bias float64, activation string, i int) {

	L.nodes[i] = &Node{
		weights:      mat.NewVecDense(len(weights), weights),
		bias:         bias,
		activation:   activation,
		parent_layer: L,
		index:        i}
	L.nodes[i].Compute_node()

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
weights: %v
bias: %f
out: %f
activation_function: %s

`, n.weights, n.bias, n.Get_output(), n.activation)

}
