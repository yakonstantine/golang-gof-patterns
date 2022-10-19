package bridge

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile(text string) {
	fmt.Println("Epson prints: '", text, "'")
}
