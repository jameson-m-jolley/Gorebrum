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

// this is the structur of the layers in the network
// the goal is to apstrat all the
type Layer struct {
	output         *mat.VecDense
	nodes          []*Node
	activation     string
	parent_network *Gorebrum
	index          int
}

// getter

func (L *Layer) Get_nodes() []*Node {
	return L.nodes
}

func (L *Layer) Get_activation() string {
	return L.activation
}

// setters

// Compute_Layer
func (L Layer) Compute_Layer() {
	for i := 0; i < len(L.nodes); i++ {
		L.nodes[i].Compute_node()
	}
	Layer_Activation_Functions[L.activation](L.output, L)
}

// this is a funtion to add new nodes to the layer
func (net *Gorebrum) New_Layer(length int, inputsLen int, activation string, index int) {
	// this needs to gen a new layer and assign the val of 0 to all the values so that the model can be trainded properly

	Layer := &Layer{
		nodes:          make([]*Node, length),
		activation:     activation,
		index:          index,
		parent_network: net,
		output:         mat.NewVecDense(length, make([]float64, length)),
	}
	for i := 0; i <= length-1; i++ {
		Layer.New_Node(make([]float64, inputsLen), 0, "ReLU", i)
	}
	net.layers[index] = Layer
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
nodes: %v
activation_function: %s

`, L.Get_nodes(), L.activation)
}
