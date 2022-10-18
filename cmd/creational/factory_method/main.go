package main

import (
	"fmt"
	fm "gof/src/creational/factory_method"
)

func main() {
	docx, err := fm.CreateDocumentProcessor(fm.Docx)
	if err != nil {
		fmt.Println(err)
		return
	}

	docx.Write("Hello from Docx")
	docx.Render()

	pdf, err := fm.CreateDocumentProcessor(fm.PDF)
	if err != nil {
		fmt.Println(err)
		return
	}

	pdf.Write("Hello from PDF")
	pdf.Render()

	tmp, err := fm.CreateDocumentProcessor(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	tmp.Write("Hello from TMP")
	tmp.Render()
}
