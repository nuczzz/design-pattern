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

// GetSingletonInstance get singleton instance of thread safe method
func GetSingletonInstance() *singleton {
	once.Do(func() {
		singletonInstance = &singleton{
			id: time.Now().UnixNano(),
		}
	})

	return singletonInstance
}

// GetInstanceId get instance id to distinguish and verify instance
func (s *singleton) GetInstanceId() int64 {
	return s.id
}
