package singleton

import (
	"time"
	"sync"
)

// Singleton Pattern Definition:
// Ensure a class has only one instance,
// and provide a global point of access to it

type SingletonPattern struct {
	id int64
}

var (
	once sync.Once
	singletonInstance *SingletonPattern
)

func NewInstance() *SingletonPattern {
	once.Do(func(){
		singletonInstance = &SingletonPattern{
			id: time.Now().UnixNano(),
		}
	})

	return singletonInstance
}

func (s *SingletonPattern) GetInstanceId() int64 {
	return s.id
}
