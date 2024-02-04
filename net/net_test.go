package net

import (
	"testing"
)

// network test
func TestNewNetwork(t *testing.T) {
	// Create a new network

	for i := 2; i < 100; i++ {
		In_d := i
		Out_d := i
		Width := i
		Depth := i
		network := New_Network(In_d, Width, Depth, Out_d)

		// Verify if the network was created successfully
		if network == nil {
			t.Error("Expected a new network to be created, got nil")
		}

		// Verify if the network has the correct dimensions
		if len(network.Layers) != Width {
			t.Errorf("Expected %d layers, got %d", Width, len(network.Layers))
		}
		//test for the input array size
		if len(network.input.RawVector().Data) != In_d {
			t.Errorf("Expected input array lenght of %d , got %d", In_d, len(network.input.RawVector().Data))
		}
		//test for output array size
		if len(network.Output.RawVector().Data) != Out_d {
			t.Errorf("Expected output array lenght of %d , got %d", Out_d, len(network.Output.RawVector().Data))
		}
		//test for the

		// test for the with of the hidden layers this is the depth paramiter
		if len(network.Layers[1].Nodes) != Depth {
			t.Errorf("Expected hidden layers array lenght of %d , got %d", Depth, len(network.Output.RawVector().Data))
		}
	} //end of test loop
}

//end of net tests
//--------------------------------------------------------------------------------------------------------------

//layer test

func TestLayerComputeLayer(t *testing.T) {
	// Create a layer with some nodes
	// network := New_Network(4, 3, 2, 2)

	// Compute the layer

	// Add assertions here to check if the layer was computed correctly
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
