package facade

import (
	"fmt"
	dcmr "gof/src/structural/facade/vndr/documenter"
)

type documenterFacade struct {
	dc *dcmr.Documenter
}

func NewDocxDocument(dc *dcmr.Documenter) *documenterFacade {
	return &documenterFacade{
		dc: dc,
	}
}

func (d *documenterFacade) Write(path, str string) (err error) {
	err = d.dc.WriteTo(path, str)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return
}

func (d *documenterFacade) Render(path string, source MyObj) (err error) {
	err = dcmr.Replace(d.dc, path, source)
	if err != nil {
		return fmt.Errorf("error render the file: %w", err)
	}
	return
}
