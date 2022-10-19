package singleton

import (
	"sync"
)

var lock = &sync.Mutex{}

func GetThreadSafeSingleton() *Singleton {
	if singleInstance != nil {
		return singleInstance
	}

	lock.Lock()
	defer lock.Unlock()

	if singleInstance != nil {
		return singleInstance
	}

	singleInstance = &Singleton{}

	return singleInstance
}
