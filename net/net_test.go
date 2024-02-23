package net

import (
	"math/rand"
	"testing"
	"time"
)

// network test

func TestNewNetwork(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	// Create a new network

	for i := 2; i < 10; i++ {
		In_d := rand.Intn(9) + 2
		Out_d := rand.Intn(9) + 2
		Width := rand.Intn(9) + 3
		Depth := rand.Intn(9) + 2
		network := New_Network(In_d, Width, Depth, Out_d)

		// Verify if the network was created successfully
		if network == nil {
			network.Display_Network()
			t.Error("Expected a new network to be created, got nil")
		}

		// Verify if the network has the correct dimensions
		if len(network.Layers) != Width {
			network.Display_Network()
			t.Errorf("Expected %d layers, got %d", Width, len(network.Layers))
		}
		//test for the input array size
		if len(network.input.RawVector().Data) != In_d {
			network.Display_Network()
			t.Errorf("Expected input array lenght of %d , got %d", In_d, len(network.input.RawVector().Data))
		}
		//test for output array size
		if len(network.Output.RawVector().Data) != Out_d {
			network.Display_Network()
			t.Errorf("Expected output array lenght of %d , got %d", Out_d, len(network.Output.RawVector().Data))
		}
		//test for the

		// test for the with of the hidden layers this is the depth paramiter
		if len(network.Layers[1].Nodes) != Depth {
			t.Errorf("Expected hidden layers array lenght of %d , got %d", Depth, len(network.Layers[1].Nodes))
		}
	} //end of test loop
}

//end of net tests
//--------------------------------------------------------------------------------------------------------------

//layer test

func TestLayerComputeLayer(t *testing.T) {
	// Create a layer with some nodes with the min set up
	network := New_Network(2, 3, 2, 2)

	//grab a layer
	network.Layers[1].parent_network.Set_inputs([]float64{0, 0})

	//set the output of the last layer

}

//end of layer tests
//--------------------------------------------------------------------------------------------------------------

// node test
func TestNodeComputeNode(t *testing.T) {
	// Create a node with some inputs, weights, bias, and activation

	// Set parent layer for the node

	// Compute the node

	// Add assertions here to check if the node output is computed correctly
}

//end of node tests
//--------------------------------------------------------------------------------------------------------------
