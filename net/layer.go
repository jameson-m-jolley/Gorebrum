package net

import "gonum.org/v1/gonum/mat"

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
func (L *Layer) Get_inputs() *mat.VecDense {
	return L.inputs
}

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

//
