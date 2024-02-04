package main

import (
	"github.com/jameson-m-jolley/Gorebrum/net"
	
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/

func main() {

	//demo.gob was created with this code
	//network := net.New_Network(4, 3, 3, 4)
	//network.Set_inputs([]float64{5, 2, 8.6, 7})
	//network.Computer_fraward_pass()
	//network.Display_Network()
	//network.Encode_model("demo")

	network := net.Dcode_model("demo.gob")
	network.Set_inputs([]float64{5, 2, 8.6, 7})
	network.Computer_fraward_pass()
	network.Display_Network()
	print(&network.Output)
}

func loop__(){

//get the inputs for the nn 

//compute the forward pass

//determan if a traning pass is needed 

//conect outputs to functaionalty

}