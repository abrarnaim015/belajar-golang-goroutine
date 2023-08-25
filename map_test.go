package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

// func AddToMap(data *sync.Map, value int, group *sync.WaitGroup)  {
// 	defer group.Done()

// 	group.Add(1)
// 	data.Store(value, value)
// }

func TestMap(t *testing.T)  {
	data := sync.Map {}
	group := sync.WaitGroup {}

	addToMap := func (value int)  {
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go addToMap(i)
		// go AddToMap(&data, i, &group)
		group.Done()
	}

	group.Wait()

	data.Range(func(key, value interface {}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}