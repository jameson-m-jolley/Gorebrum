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
	loop__(net.New_trainer(network))
}

// this  was uses to train the network
func train_imgs(coach *net.Trainer, hotvet int, path string) {

}
