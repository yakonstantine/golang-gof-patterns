package main

import (
	"fmt"
	"gof/src/structural/composite"
)

func main() {
	f1 := composite.NewFolder("root")
	f2 := composite.NewFolder("usr")
	f3 := composite.NewFolder("tmp")
	f4 := composite.NewFolder("fooBar")

	f1.Add(f2)
	f1.Add(f3)
	f2.Add(f4)

	fl1 := composite.NewFile("text.txt")
	fl2 := composite.NewFile("bar.txt")
	fl3 := composite.NewFile("foo.txt")
	fl4 := composite.NewFile("word.txt")
	fl5 := composite.NewFile("key.txt")
	fl6 := composite.NewFile("note.txt")

	f1.Add(fl1)
	f2.Add(fl2)
	f2.Add(fl3)
	f4.Add(fl4)
	f4.Add(fl5)
	f4.Add(fl6)

	name, ok := f1.Search("bar")

	fmt.Println("Result is:", ok, "File name:", name)
}
