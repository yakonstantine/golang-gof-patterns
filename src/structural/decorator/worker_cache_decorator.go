package decorator

import "fmt"

type cachedWorker struct {
	decorated Worker
	cache     map[string]string
}

func NewCachedWorker(w Worker) *cachedWorker {
	return &cachedWorker{
		decorated: w,
		cache:     map[string]string{},
	}
}

func (cw *cachedWorker) DoWork(command string) (res string, err error) {
	if v, ok := cw.cache[command]; ok {
		res = fmt.Sprintf("from cache: %s", v)
		return
	}
	res, err = cw.decorated.DoWork(command)
	cw.cache[command] = res
	return
}
