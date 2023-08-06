package main

import (
	"github.com/jameson-m-jolley/Gorebrum/net"
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/

func main() {
	network := net.New_Network(3, 3, 3, 2)
	network.Display_Network()
	network.Computer_fraward_pass()

}
