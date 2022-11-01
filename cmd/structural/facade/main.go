package main

import (
	"gof/src/structural/facade"
	"gof/src/structural/facade/vndr/documenter"
)

func main() {
	dcmr := &documenter.Documenter{}
	df := facade.NewDocxDocument(dcmr)
	df.Write("my_file.txt", "hello, world!")
	df.Render("my_file.txt", facade.MyObj{Name: "Foo", Age: 18})
}
