package main

import "gof/src/structural/bridge"

func main() {
	eps := bridge.Epson{}
	hp := bridge.HP{}

	pc := bridge.NewPC(&eps)
	mac := bridge.NewMac(&hp)

	pc.Print("Hi from PC")
	mac.Print("Hi from Mac")
}
