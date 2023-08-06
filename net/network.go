package net

import (
	"fmt"

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

// getters for the network
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

// this is the functon that makes a new network
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
			network.New_Layer(depth, in_dimensions, "ReLU", i)
		} else if i == width-1 {
			network.New_Layer(out_dimensions, depth, "ReLU", i)
		} else {
			network.New_Layer(depth, depth, "ReLU", i)
		}
	}
	return network
}

// the farward pass for the network
func (net *Gorebrum) Computer_fraward_pass() {
	for i := 0; i < net.width-1; i++ {
		net.layers[i].Compute_Layer()
	}
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
