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
	inputs       *mat.VecDense
	weights      *mat.VecDense
	bias         float64
	activation   string
	output       float64
	parent_layer *Layer
}

// geters

func (n Node) Get_inputs() *mat.VecDense {
	return n.inputs
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

// setters
func (n *Node) Set_inputs(inputs *mat.VecDense) {
	n.inputs = inputs
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
}

// this sets the node inside a layer at the index if intdex>len(Layer.nodes) the func will error
// use at your one risk this can brake the NN
func (n *Node) Set_parent_layer(Layer *Layer, index int) {
	Layer.nodes[index] = n
}

// method for computing the output of a node
func (n Node) Compute_node(inputs *mat.VecDense) {
	//sets the inputs of the node
	n.Set_inputs(inputs)
	// hash of actavation functions sored in the ./activation.go file this is done to sapport diffrent actavation functions throught out the network of layer level
	n.Set_output(Node_Activation_Functions[n.activation](mat.Dot(n.inputs, n.weights)+n.bias, n))
}

// makes a new node
func New_Node(weights []float64, bias float64, activation string, parent_layer *Layer) *Node {
	return &Node{
		weights:      mat.NewVecDense(len(weights), weights),
		bias:         bias,
		activation:   activation,
		parent_layer: parent_layer}
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

`, n.inputs, n.weights, n.bias, n.Get_output(), n.activation)

}
