package factorymethod

import "fmt"

type PdfDocument struct {
	text string
}

func NewPdfDocument() *PdfDocument {
	return &PdfDocument{}
}

func (doc *PdfDocument) Write(str string) {
	doc.text = str
}

func (doc *PdfDocument) Render() {
	fmt.Printf("PDF text is '%s'", doc.text)
	fmt.Println()
}
