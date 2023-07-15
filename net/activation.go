package net

/*
▄▀█ █▀▀ ▀█▀ █ █░█ ▄▀█ ▀█▀ █ █▀█ █▄░█   █▀▀ █░█ █▄░█ █▀▀ ▀█▀ █ █▀█ █▄░█ █▀
█▀█ █▄▄ ░█░ █ ▀▄▀ █▀█ ░█░ █ █▄█ █░▀█   █▀░ █▄█ █░▀█ █▄▄ ░█░ █ █▄█ █░▀█ ▄█

we are using a dict to store the funtions. each function will be stored under a key that is a string so that each node can have a key to a actavation function
*/
var Activation_functions map[string]func(float64) float64 = map[string]func(float64) float64{
	"ReLU": ReLU,
}

// rectified linear: takes a float64 as var x and returns 0 if input < 0 else returns var x
func ReLU(x float64) float64 {
	if x < 0 {
		return 0
	} else {
		return x
	}
}
