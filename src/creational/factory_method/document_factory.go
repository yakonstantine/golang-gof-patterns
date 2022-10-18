package factorymethod

import "fmt"

func CreateDocumentProcessor(dt DocType) (DocumentProcessor, error) {
	switch dt {
	case Docx:
		return NewDocxDocument(), nil
	case PDF:
		return NewPdfDocument(), nil
	default:
		return nil, fmt.Errorf("unsupported document type '%d'", dt)
	}
}
