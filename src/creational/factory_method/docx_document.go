package factorymethod

import "fmt"

type DocxDocument struct {
	text string
}

func NewDocxDocument() *DocxDocument {
	return &DocxDocument{}
}

func (doc *DocxDocument) Write(str string) {
	doc.text = str
}

func (doc *DocxDocument) Render() {
	fmt.Printf("Docx text is '%s'", doc.text)
	fmt.Println()
}
