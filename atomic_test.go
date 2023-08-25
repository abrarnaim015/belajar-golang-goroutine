package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T)  {
	wg := sync.WaitGroup {}
	var x int64 = 0
	// x := 0
	
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}
	
	wg.Wait()
	fmt.Println("Counter : ", x)
}