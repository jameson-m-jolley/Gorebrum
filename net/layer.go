package net

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

/*
██╗░░░░░░█████╗░██╗░░░██╗███████╗██████╗░
██║░░░░░██╔══██╗╚██╗░██╔╝██╔════╝██╔══██╗
██║░░░░░███████║░╚████╔╝░█████╗░░██████╔╝
██║░░░░░██╔══██║░░╚██╔╝░░██╔══╝░░██╔══██╗
███████╗██║░░██║░░░██║░░░███████╗██║░░██║
╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝
*/

type Layer struct {
	inputs     *mat.VecDense
	nodes      []*Node
	output     *mat.VecDense
	activation string
}

// getter

func (L *Layer) Get_nodes() []*Node {
	return L.nodes
}

func (L *Layer) Get_output() *mat.VecDense {
	return L.output
}

func (L *Layer) Get_activation() string {
	return L.activation
}

// setters

func (L *Layer) Set_inputs(inputs *mat.VecDense) {
	L.inputs = inputs
}

//other

// Compute_Layer
func (L Layer) Compute_Layer() *mat.VecDense {
	for i := 0; i < len(L.nodes); i++ {
		L.output.SetVec(i, L.nodes[i].Compute_node(L.inputs))
	}
	return Layer_Activation_Functions[L.activation](L.output, L)
}

// this is a funtion to add new nodes to the layer
func New_Layer(length int, activation string) *Layer {
	// this needs to gen a new layer and assign the val of 0 to all the values so that the model can be trainded properly
	// buffer for the nodes that need to be made
	node_buff := make([]*Node, 0, length)

	for i := 0; i <= length; i++ {
		node_buff = append(node_buff, Newnode(make([]float64, length), 0, activation))
	}
	return &Layer{nodes: node_buff, activation: activation}
}

// display the info for a layer with
func (L Layer) Display_info() {
	fmt.Printf(`
---------------------------------------------------------------------

██╗░░░░░░█████╗░██╗░░░██╗███████╗██████╗░
██║░░░░░██╔══██╗╚██╗░██╔╝██╔════╝██╔══██╗
██║░░░░░███████║░╚████╔╝░█████╗░░██████╔╝
██║░░░░░██╔══██║░░╚██╔╝░░██╔══╝░░██╔══██╗
███████╗██║░░██║░░░██║░░░███████╗██║░░██║
in: %v
nodes: %v
out: %v
activation_function: %s

`, L.inputs, L.Get_nodes(), L.Get_output(), L.activation)
}
