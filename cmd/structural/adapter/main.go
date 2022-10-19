package main

import "gof/src/structural/adapter"

func main() {
	s := adapter.EmailSender{}
	cl1 := adapter.NewClient(&s)
	cl1.Do()

	so := adapter.OldSender{}
	a := adapter.NewOldSenderAdapter(&so)
	cl2 := adapter.NewClient(a)
	cl2.Do()
}
