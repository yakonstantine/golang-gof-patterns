package bridge

type PC struct {
	prt Printer
}

func NewPC(p Printer) *PC {
	return &PC{prt: p}
}

func (pc *PC) Print(text string) {
	pc.prt.PrintFile(text)
}
