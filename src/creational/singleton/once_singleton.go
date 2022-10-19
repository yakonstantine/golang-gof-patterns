package singleton

import (
	"sync"
)

var once sync.Once

func GetOnceSingletonInstance() *Singleton {
	if singleInstance != nil {
		return singleInstance
	}

	once.Do(
		func() {
			singleInstance = &Singleton{}
		})

	return singleInstance
}
