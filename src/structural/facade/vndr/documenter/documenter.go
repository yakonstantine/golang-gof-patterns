package documenter

import "fmt"

type Documenter struct {
}

func (d *Documenter) Read(docPath string) (str string, err error) {
	err = d.open(docPath)
	if err != nil {
		return
	}
	defer d.close(docPath)
	str = fmt.Sprintf("Read from:'%s'", docPath)
	return
}

func (d *Documenter) WriteTo(docPath, text string) (err error) {
	err = d.open(docPath)
	if err != nil {
		return
	}
	defer d.close(docPath)
	fmt.Printf("Write to:'%s', text:'%s'\n", docPath, text)
	return
}

func Replace[T any](d *Documenter, docPath string, source T) (err error) {
	err = d.open(docPath)
	if err != nil {
		return
	}
	defer d.close(docPath)
	fmt.Printf("Replace in:'%s', source:'%v'\n", docPath, source)
	return
}

func (d *Documenter) open(docPath string) error {
	fmt.Printf("Open doc:'%s'\n", docPath)
	return nil
}

func (d *Documenter) close(docPath string) error {
	fmt.Printf("Close doc:'%s'\n", docPath)
	return nil
}
