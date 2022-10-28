package main

import (
	"fmt"
	"gof/src/structural/decorator"
)

func main() {
	w := getWorker()
	res, _ := w.DoWork("write")
	fmt.Println(res)
	res, _ = w.DoWork("read")
	fmt.Println(res)
	res, _ = w.DoWork("write")
	fmt.Println(res)
	res, _ = w.DoWork("read")
	fmt.Println(res)
}

func getWorker() decorator.Worker {
	w := decorator.NewEmployee("John")
	cw := decorator.NewCachedWorker(w)
	lw := decorator.NewLogWorker(cw)
	return lw
}
