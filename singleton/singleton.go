package singleton

import (
	"sync"
	"time"
)

// Singleton Pattern Definition:
// Ensure a class has only one instance,
// and provide a global point of access to it.

type Singleton struct {
	id int64
}

var (
	once              sync.Once
	singletonInstance *Singleton
)

func GetSingletonInstance() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{
			id: time.Now().UnixNano(),
		}
	})

	return singletonInstance
}

func (s *Singleton) GetInstanceId() int64 {
	return s.id
}
