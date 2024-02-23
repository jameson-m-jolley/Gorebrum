package net

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/gonum/mat"
)

type Gorebrum struct {
	input          *mat.VecDense
	Layers         []*Layer
	Output         *mat.VecDense
	Width          int
	Depth          int
	Out_dimensions int
	In_dimensions  int
	Log            []string // for training
}

// getters
// this is done because I want the option in the future to control how models are
// interact with at a base level
// I do not intend to use dot notation to set and extract variables i aslo do not want
// the fields of the structs to be visible to external programs

func (N Gorebrum) Get_Output() *mat.VecDense {
	return N.Output
}
func (N Gorebrum) Get_input() *mat.VecDense {
	return N.input
}
func (N Gorebrum) Get_Layers() []*Layer {
	return N.Layers
}
func (N Gorebrum) Get_layer(index int) *Layer {
	return N.Layers[index]
}

// setters for the network
func (N *Gorebrum) Set_Layer(index int, layer *Layer) {
	N.Layers[index] = layer
}

// sets an input to the seme vec that is passed to it
func (N *Gorebrum) Set_inputs(inputs []float64) {
	if N.In_dimensions == len(inputs) {
		N.input = mat.NewVecDense(N.In_dimensions, inputs)
	} else {
		panic("length of inputs is out of range len(inputs) must be == N.In_dimensions")
	}

}

/*
this is the functon that makes a new network,
New_Network takes 4 variabels as int as fallows

In_dimensions int, (this is the length of the vector that is pasesed into the network)

Width int, (the number of Layers includeing the frist and last Layers)

Depth int, (the lenght of the vectors in the hidden Layers)

Out_dimensions int, (the length of the Output layer)

the function will then return a *Gorebrum:

		type Gorebrum struct {
		input          *mat.VecDense
		Layers         []*Layer
		Output         *mat.VecDense
		Width          int
		Depth          int
		Out_dimensions int
		In_dimensions  int
	}

	the Output layer will allwase be a softmax actavation
*/
func New_Network(
	In_dimensions int,
	Width int,
	Depth int,
	Out_dimensions int) *Gorebrum {
	//creates a new network and creates the Layers and Nodes
	network := &Gorebrum{
		In_dimensions:  In_dimensions,
		Out_dimensions: Out_dimensions,
		Width:          Width,
		Depth:          Depth,
		Layers:         make([]*Layer, Width),
		input:          mat.NewVecDense(In_dimensions, make([]float64, In_dimensions)),
		Output:         mat.NewVecDense(Out_dimensions, make([]float64, Out_dimensions)),
	}
	for i := 0; i <= Width-1; i++ {
		if i == 0 {
			// frist layer
			network.New_Layer(Depth, In_dimensions, "NA", i)
		} else if i == Width-1 {
			// Output layer
			network.New_Layer(Out_dimensions, Depth, "SoftMax", i)
		} else {
			// hidden Layers
			network.New_Layer(Depth, Depth, "NA", i)
		}
	}
	return network
}

// the farward pass for the network
func (N *Gorebrum) Computer_fraward_pass() {
	for i := 0; i < N.Width; i++ {
		N.Layers[i].Compute_Layer()
	}
	N.Update_Output()
}

// this updates the Output for the network
func (N *Gorebrum) Update_Output() {
	N.Output = N.Get_layer(N.Width - 1).Output
}

// this displays the network in all its glory and wonder
func (N *Gorebrum) ToXML() string {
	XML := "<NN class =\"model\">"
	for i := 0; i < len(N.Layers); i++ {
		XML += N.Layers[i].ToXML()
	}
	XML += "</NN>"
	return XML

}

func (N *Gorebrum) Display_Network() {
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
Layers         :%v
Output         :%v
Width          :%v
Depth          :%v
Out_dimensions :%v
In_dimensions  :%v
`, N.input, N.Layers, N.Output, N.Width, N.Depth, N.Out_dimensions, N.In_dimensions)

	if N.Depth <= 50 {
		nns := "_______________________Network__________________________\n"
		for i := 0; i < len(N.Layers); i++ {
			nns += fmt.Sprint(i) + ":"
			for b := 0; b < len(N.Layers[i].Nodes); b++ {
				nns += "|" + fmt.Sprint(N.Layers[i].Nodes[b].Output) + "|"
			}
			nns += "\n"
		}
		fmt.Print(nns)
	}
}

// encodes the network and wrigths the model to a gob file wich
func (N *Gorebrum) Encode_model(name string) {

	//makes a filefore
	f, err := os.Create(name + ".gob")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	code := gob.NewEncoder(f)
	//encodes and logs err
	if err := code.Encode(N); err != nil {
		log.Fatal(err)
	}

}

func Dcode_model(name string) *Gorebrum {

	//tests to see if the file is their
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	code := gob.NewDecoder(f)
	//encodes and logs err
	var N Gorebrum
	if err := code.Decode(&N); err != nil {
		log.Fatal(err)
	}

	N.repare()

	return &N
}

// this is a 2 part recusive function that inclues repare_nodes
func (N *Gorebrum) repare() {
	for i := 0; i <= len(N.Layers)-1; i++ {
		N.Layers[i].Set_parent_network(N)
		for f := 0; f <= len(N.Layers[i].Nodes)-1; f++ {
			N.Layers[i].Nodes[f].Set_parent_layer(N.Layers[i])
		}
	}
}
