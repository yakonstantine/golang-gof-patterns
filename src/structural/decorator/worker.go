package decorator

type Worker interface {
	DoWork(command string) (res string, err error)
}
