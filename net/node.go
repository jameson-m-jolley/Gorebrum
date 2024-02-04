package net

import (
	"fmt"
	"math/rand"

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

// this is the struct for creating Nodes
// this will eventualy be a nural network
type Node struct {
	Weights      *mat.VecDense
	inputs       *mat.VecDense // the field is kept private because it is not needed in the encodeing
	Bias         float64
	Activation   string
	Output       float64
	parent_layer *Layer // remove this from the encodeing and fix it in the decode file
	Index        int
}

// geters
// this is done because I want the option in the future to control how models are
// interact with at a base level
// I do not intend to use dot notation to set and extract variables i aslo do not want
// the fields of the structs to be visible to external programs

// this function grabs the Outputs from the layer before it.
// unless it is the frist layer, in that case the input for the network is used
func (n Node) Get_input() *mat.VecDense {
	if n.parent_layer.Index == 0 {
		n.inputs = n.parent_layer.parent_network.input
		return n.parent_layer.parent_network.input
	} else {
		n.inputs = n.parent_layer.parent_network.Layers[n.parent_layer.Index-1].Output
		return n.parent_layer.parent_network.Layers[n.parent_layer.Index-1].Output
	}
}

// returns the weights of the node as a *mat.VecDence
func (n Node) Get_Weights() *mat.VecDense {
	return n.Weights
}

// returns the bias of the node
func (n Node) Get_Bias() float64 {
	return n.Bias
}

// reternds the activation of the network.
func (n Node) Get_Activation() string {
	return n.Activation
}

// returns the output of the node
func (n Node) Get_output() float64 {
	return n.Output
}

// returs the address of the layer that the node is a part of
func (n Node) Get_parent_layer() *Layer {
	return n.parent_layer
}

// sets the weights of the node to a *mat.VecDence
func (n *Node) Set_Weights(Weights *mat.VecDense) {
	n.Weights = Weights
}

// sets the Bias of a node to a float64
func (n *Node) Set_Bias(Bias float64) {
	n.Bias = Bias
}

func (n *Node) Set_Activation(Activation string) {
	n.Activation = Activation
}

func (n *Node) Set_input(input *mat.VecDense) {
	n.inputs = input
}

func (n *Node) Set_Output(Output float64) {
	n.Output = Output
	n.parent_layer.Output.SetVec(n.Index, n.Output)
}

func (n *Node) Set_parent_layer(Layer *Layer) {
	n.parent_layer = Layer
}

// this sets the val of each poit in the vec to be a rand.Float64
// this only needs to be called at the creation of the NN
func (n *Node) scrambler(i int) {
	if i < n.Weights.Len() {
		n.Weights.SetVec(i, rand.Float64())
		i++
		go n.scrambler(i)
	} else {
		n.Set_Bias(rand.Float64())
	}
}

// method for computing the Output of a node
func (n *Node) Compute_node() {
	//updates the inputs
	// hash of actavation functions sored in the ./Activation.go file this is done to sapport diffrent actavation functions throught out the network of layer level
	n.Set_Output(Node_Activation_Functions[n.Activation](mat.Dot(n.Get_input(), n.Get_Weights())+n.Bias, n))
}

// makes a new node

func (L *Layer) New_Node(Weights []float64, Bias float64, Activation string, i int) {
	//makes the new node in mem
	node := &Node{
		Weights:      mat.NewVecDense(len(Weights), Weights),
		Bias:         Bias,
		Activation:   Activation,
		parent_layer: L,
		Index:        i}
	// if the len of the arry is less then the index of the arr we need to append  the new e
	if len(L.Nodes) <= i {
		L.Nodes = append(L.Nodes, node)
	} else {
		L.Nodes[i] = node
	}
	L.Nodes[i].Compute_node()
	L.Nodes[i].scrambler(0)

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
Weights: %v
Bias: %f
out: %f
Activation_function: %s

`, n.Weights, n.Bias, n.Get_output(), n.Activation)

}
