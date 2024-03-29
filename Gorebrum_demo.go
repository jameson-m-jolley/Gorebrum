package main

import (
	"github.com/jameson-m-jolley/Gorebrum/net"
)

// banner\output art wich was genarated using https://fsymbols.com/generators/carty/

func main() {

	//demo.gob was created with this code
	//network := net.New_Network(10000, 6, 6, 3)
	//network.Encode_model("demo")

	network := net.Dcode_model("demo.gob")
	//network.Set_inputs([]float64{5, 2, 8.6, 7})
	//network.Computer_fraward_pass()
	network.Display_Network()
	trainer := net.New_trainer(network)
	trainer.Set_algorithm("rand_mutation")
	trainer.Compute_Training_Pass()

}
