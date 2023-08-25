package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T)  {
	t.Skip()
	wg := sync.WaitGroup {}
	ticker := time.NewTicker(1 * time.Second)

	wg.Add(1)
	go func() {
		time.AfterFunc(5 * time.Second, func() {
			defer wg.Done()
			ticker.Stop()
		})
		// time.Sleep(5 * time.Second)
		// ticker.Stop()
	}()
	
	for time := range ticker.C {
		fmt.Println(time)
	}
	
	wg.Wait()
}

func TestTick(t *testing.T)  {
	channel := time.Tick(1 * time.Second)
	
	for time := range channel {
		fmt.Println(time)
	}
	
}