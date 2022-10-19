package main

import (
	"fmt"
	"gof/src/creational/singleton"
)

func main() {
	si := singleton.GetThreadSafeSingleton()
	si.Value = 123

	fmt.Println(singleton.GetOnceSingletonInstance().Value)
	fmt.Println(singleton.GetThreadSafeSingleton().Value)
}
