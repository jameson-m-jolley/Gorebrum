package net

//this will be the implamentation of traing the network

type trainer struct {
	inputs    [][]float64
	outputs   [][]float64
	network   *Gorebrum
	hotvector int
	loss      float64
}

// this creates a new training obj for training a model on a data set
// doing it this way will give is freedom to implement different training strategies
func New_trainer(network *Gorebrum) *trainer {
	return &trainer{network: network}
}
