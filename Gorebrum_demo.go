package main

import (
	"github.com/jameson-m-jolley/Gorebrum/net"
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/

func main() {
	network := net.New_Network(3, 3, 3, 2)
	network.Set_inputs([]float64{5, 2, 8.6})
	network.Computer_fraward_pass()
	network.Display_Network()

	println(network.Compute_loss(0))
}
