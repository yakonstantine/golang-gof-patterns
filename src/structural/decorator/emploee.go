package decorator

import "fmt"

type employee struct {
	name string
}

func NewEmployee(name string) *employee {
	if name == "" {
		name = "Stranger"
	}
	return &employee{
		name: name,
	}
}

func (e *employee) DoWork(command string) (string, error) {
	return fmt.Sprintf("%s is working on %s", e.name, command), nil
}

func (e *employee) Name() string {
	return e.name
}
