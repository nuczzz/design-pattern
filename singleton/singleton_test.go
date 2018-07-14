package singleton

import (
	"testing"
	"fmt"
	"sync"
)

func TestNewInstance(t *testing.T) {
	var wg sync.WaitGroup
	var goroutineNum = 5

	wg.Add(goroutineNum)
	for index := 0; index < goroutineNum; index ++ {
		go func(index int){
			singleton := NewInstance()
			fmt.Printf("goroutine[%v]->singletonId[%v]\n", index, singleton.GetInstanceId())
			wg.Done()
		}(index)
	}
	wg.Wait()
}
