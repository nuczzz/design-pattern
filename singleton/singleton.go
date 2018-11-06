package singleton

import (
	"sync"
	"time"
)

// Singleton Pattern Definition:
// Ensure a class has only one instance,
// and provide a global point of access to it.

type singleton struct {
	id int64
}

var (
	once              sync.Once
	singletonInstance *singleton
)

func GetSingletonInstance() *singleton {
	once.Do(func() {
		singletonInstance = &singleton{
			id: time.Now().UnixNano(),
		}
	})

	return singletonInstance
}

func (s *singleton) GetInstanceId() int64 {
	return s.id
}
