package bridge

type Mac struct {
	prt Printer
}

func NewMac(p Printer) *Mac {
	return &Mac{prt: p}
}

func (m *Mac) Print(text string) {
	m.prt.PrintFile(text)
}
