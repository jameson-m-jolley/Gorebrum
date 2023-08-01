package main

import (
	"fmt"

	"github.com/jameson-m-jolley/Gorebrum/net"
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/
func print_logo() {
	fmt.Print(`

░██████╗░░█████╗░██████╗░███████╗██████╗░██████╗░██╗░░░██╗███╗░░░███╗
██╔════╝░██╔══██╗██╔══██╗██╔════╝██╔══██╗██╔══██╗██║░░░██║████╗░████║
██║░░██╗░██║░░██║██████╔╝█████╗░░██████╦╝██████╔╝██║░░░██║██╔████╔██║
██║░░╚██╗██║░░██║██╔══██╗██╔══╝░░██╔══██╗██╔══██╗██║░░░██║██║╚██╔╝██║
╚██████╔╝╚█████╔╝██║░░██║███████╗██████╦╝██║░░██║╚██████╔╝██║░╚═╝░██║
░╚═════╝░░╚════╝░╚═╝░░╚═╝╚══════╝╚═════╝░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░╚═╝
Created by: Jameson Jolley
version: 1.0
`)
}

func main() {

	print_logo()
	anode := net.Newnode([]float64{1.8, 8.3, 7.1}, 4.5, "Log")
	anode.Display_info()

	alayer := net.New_Layer(5, 3, "ReLU")
	alayer.Display_info()

	alayer.Get_nodes()[2].Display_info()

	network := net.NewNetwork(5, 5, 5, 2)
	fmt.Print(network)
	network.Display_Network()

}
