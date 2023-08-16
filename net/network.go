package net

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

type Gorebrum struct {
	input          *mat.VecDense
	layers         []*Layer
	output         *mat.VecDense
	width          int
	depth          int
	out_dimensions int
	in_dimensions  int
}

// getters
// this is done because I want the option in the future to control how models are
// interact with at a base level
// I do not intend to use dot notation to set and extract variables i aslo do not want
// the fields of the structs to be visible to external programs

func (net Gorebrum) Get_output() *mat.VecDense {
	return net.output
}
func (net Gorebrum) Get_input() *mat.VecDense {
	return net.input
}
func (net Gorebrum) Get_layers() []*Layer {
	return net.layers
}
func (net Gorebrum) Get_layer(index int) *Layer {
	return net.layers[index]
}

// setters for the network
func (net *Gorebrum) Set_Layer(index int, layer *Layer) {
	net.layers[index] = layer
}

// sets an input to the seme vec that is passed to it
func (net *Gorebrum) Set_inputs(inputs []float64) {
	if net.in_dimensions == len(inputs) {
		net.input = mat.NewVecDense(net.in_dimensions, inputs)
	} else {
		panic("length of inputs is out of range len(inputs) must be == net.in_dimensions")
	}

}

/*
this is the functon that makes a new network,
New_Network takes 4 variabels as int as fallows

	in_dimensions int, (this is the length of the vector that is pasesed into the network)
	width int, (the number of layers includeing the frist and last layers)
	depth int, (the lenght of the vectors in the hidden layers)
	out_dimensions int, (the length of the output layer)

	the function will then return a *Gorebrum
	the output layer will allwase be a softmax actavation
*/
func New_Network(
	in_dimensions int,
	width int,
	depth int,
	out_dimensions int) *Gorebrum {
	//creates a new network and creates the layers and nodes
	network := &Gorebrum{
		in_dimensions:  in_dimensions,
		out_dimensions: out_dimensions,
		width:          width,
		depth:          depth,
		layers:         make([]*Layer, width),
		input:          mat.NewVecDense(in_dimensions, make([]float64, in_dimensions)),
		output:         mat.NewVecDense(out_dimensions, make([]float64, out_dimensions)),
	}
	for i := 0; i <= width-1; i++ {
		if i == 0 {
			// frist layer
			network.New_Layer(depth, in_dimensions, "NA", i)
		} else if i == width-1 {
			// output layer
			network.New_Layer(out_dimensions, depth, "SoftMax", i)
		} else {
			// hidden layers
			network.New_Layer(depth, depth, "NA", i)
		}
	}
	return network
}

// the farward pass for the network
func (net *Gorebrum) Computer_fraward_pass() {
	for i := 0; i < net.width; i++ {
		net.layers[i].Compute_Layer()
	}
	net.Update_output()
}

// this updates the output for the network
func (net *Gorebrum) Update_output() {
	net.output = net.Get_layer(net.width - 1).output
}

// -log loss
// computes the cross entropy loss at at the target vector
func (net *Gorebrum) Compute_loss(vec int) float64 {
	return -(math.Log(net.output.AtVec(vec)))

}

// this displays the network in all its glory and wonder

func (net *Gorebrum) Display_Network() {
	fmt.Printf(`

░██████╗░░█████╗░██████╗░███████╗██████╗░██████╗░██╗░░░██╗███╗░░░███╗
██╔════╝░██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██║░░░██║████╗░████║
██║░░██╗░██║░░██║██████╔╝█████╗░░██████╦╝██████╔╝██║░░░██║██╔████╔██║
██║░░╚██╗██║░░██║██╔══██╗██╔══╝░░██╔══██╗██╔══██╗██║░░░██║██║╚██╔╝██║
╚██████╔╝╚█████╔╝██║░░██║███████╗██████╦╝██║░░██║╚██████╔╝██║░╚═╝░██║
░╚═════╝░░╚════╝░╚═╝░░╚═╝╚══════╝╚═════╝░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░╚═╝
Created by: Jameson Jolley
version: 1.0
input          :%v
layers         :%v
output         :%v
width          :%v
depth          :%v
out_dimensions :%v
in_dimensions  :%v
`, net.input, net.layers, net.output, net.width, net.depth, net.out_dimensions, net.in_dimensions)

	if net.depth <= 50 {
		nns := "_______________________Network__________________________\n"
		for i := 0; i < len(net.layers); i++ {
			nns += fmt.Sprint(i) + ":"
			for b := 0; b < len(net.layers[i].nodes); b++ {
				nns += "|" + fmt.Sprint(net.layers[i].nodes[b].output) + "|"
			}
			nns += "\n"
		}
		fmt.Print(nns)
	}
}
