package singleton

import (
	"sync"
)

type Singleton struct {
	Number      int
	Description string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
