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

// setters for the network

// this is the functon that makes a new network
func NewNetwork(
	in_dimensions int,
	width int,
	depth int,
	out_dimensions int) *Gorebrum {

	//buffer fore the layers
	Layer_buff := make([]*Layer, 0, depth)
	for i := 1; i <= width; i++ {
		Layer_buff = append(Layer_buff, New_Layer(depth, depth, "ReLU"))
	}
	// sets the frist and last layers the the right sape
	Layer_buff[0] = New_Layer(depth, in_dimensions, "ReLU")
	Layer_buff[width-1] = New_Layer(out_dimensions, depth, "ReLU")

	return &Gorebrum{in_dimensions: in_dimensions, out_dimensions: out_dimensions, width: width, depth: depth, layers: Layer_buff}

}

// the farward pass for the network
func (net *Gorebrum) Computer_fraward_pass() {

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
}
