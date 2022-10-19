package bridge

import "fmt"

type HP struct {
}

func (p *HP) PrintFile(text string) {
	fmt.Println("HP prints: '", text, "'")
}
