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
	anode := net.Newnode([]float64{1.1, 2.3, 5.1}, []float64{1.8, 8.3, 7.1}, 4.5, "ReLU")
	anode.Display_info()
}
