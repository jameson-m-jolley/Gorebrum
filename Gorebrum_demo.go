package main

import (
	"github.com/jameson-m-jolley/Gorebrum/net"
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/

func main() {
	network := net.New_Network(3, 3, 3, 2)
	network.Display_Network()
	network.Set_inputs([]float64{5, 2, 8.6})
	network.Computer_fraward_pass()
	network.Display_Network()
	network.Get_layer(1).Get_nodes()[1].Display_info()

	network.Set_inputs([]float64{1, 5, 6})
	network.Computer_fraward_pass()
	network.Display_Network()
}
