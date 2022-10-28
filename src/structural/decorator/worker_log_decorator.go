package decorator

import "fmt"

type logWorker struct {
	decorated Worker
}

func NewLogWorker(w Worker) *logWorker {
	return &logWorker{
		decorated: w,
	}
}

func (lw *logWorker) DoWork(command string) (res string, err error) {
	fmt.Printf("command '%s' is started\n", command)
	res, err = lw.decorated.DoWork(command)
	if err != nil {
		fmt.Printf("error executing command '%s': %s\n", command, err)
		return
	}
	fmt.Printf("command '%s' is completed\n", command)
	return
}
