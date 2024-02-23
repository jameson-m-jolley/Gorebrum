package net

import (
	"fmt"
	"strconv"

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

// this is the structur of the Layers in the network
// the goal is to apstrat all the
type Layer struct {
	Output         *mat.VecDense
	Nodes          []*Node
	Activation     string
	parent_network *Gorebrum // remove this from the encodeing and fix it in the decode file
	Index          int
	Log            []string // for training
}

// getter
// this is done because I want the option in the future to control how models are
// interact with at a base level
// I do not intend to use dot notation to set and extract variables i aslo do not want
// the fields of the structs to be visible to external programs

func (L *Layer) Get_Nodes() []*Node {
	return L.Nodes
}

func (L *Layer) Get_parent_network() *Gorebrum {
	return L.parent_network
}

func (L *Layer) Get_Activation() string {
	return L.Activation
}

func (L *Layer) Get_Index() int {
	return L.Index
}

// setters
func (L *Layer) Set_parent_network(parent_network *Gorebrum) {
	L.parent_network = parent_network
}

func (L *Layer) Set_Index(index int) {
	L.Index = index
}

func (L *Layer) Set_Nodes(nodes []*Node) {
	L.Nodes = nodes
}

func (L *Layer) Set_Activation(activation string) {
	L.Activation = activation
}

func (L *Layer) Set_Output(output *mat.VecDense) {
	L.Output = output
}

// Compute_Layer
func (L Layer) Compute_Layer() {
	for i := 0; i < len(L.Nodes); i++ {
		L.Nodes[i].Compute_node()
	}
	Layer_Activation_Functions[L.Activation](L.Output, L)
}

// this is a funtion to add new Nodes to the layer
func (net *Gorebrum) New_Layer(length int, inputsLen int, Activation string, Index int) {
	// this needs to gen a new layer and assign the val of 0 to all the values so that the model can be trainded properly

	Layer := &Layer{
		Nodes:          make([]*Node, length),
		Activation:     Activation,
		Index:          Index,
		parent_network: net,
		Output:         mat.NewVecDense(length, make([]float64, length)),
	}
	for i := 0; i <= length-1; i++ {
		Layer.New_Node(make([]float64, inputsLen), 0, "ReLU", i)
	}
	net.Layers[Index] = Layer
}

// display the info for a layer with

func (L Layer) ToXML() string {
	XML := "<layer id=\"" + strconv.Itoa(L.Index) + "\" >"
	for i := 0; i < len(L.Nodes); i++ {
		XML += L.Nodes[i].ToXML()
	}
	XML += "</layer>"
	return XML
}

func (L Layer) Display_info() {
	fmt.Printf(`
---------------------------------------------------------------------

██╗░░░░░░█████╗░██╗░░░██╗███████╗██████╗░              
██║░░░░░██╔══██╗╚██╗░██╔╝██╔════╝██╔══██╗
██║░░░░░███████║░╚████╔╝░█████╗░░██████╔╝
██║░░░░░██╔══██║░░╚██╔╝░░██╔══╝░░██╔══██╗
███████╗██║░░██║░░░██║░░░███████╗██║░░██║
Nodes: %v
Activation_function: %s
Outputs: %v

`, L.Get_Nodes(), L.Activation, L.Output)
}
