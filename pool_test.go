package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T)  {
	pool := sync.Pool {
		New: func () interface {}  {
			return "New"
		},
	}

	group := sync.WaitGroup {}

	pool.Put("Abrar")
	pool.Put("Naim")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	// time.Sleep(3 * time.Second)
	group.Wait()
	fmt.Println("Done")
}