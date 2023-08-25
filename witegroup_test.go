package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

// untuk penggunan await dan jangan lupa di done()

func RunAsynchronous(grup *sync.WaitGroup, i int)  {
	defer grup.Done()

	grup.Add(1)

	fmt.Println("Hello ", i)
	// time.Sleep(1 * time.Second)
}

func TestWitGrup(t *testing.T)  {
	grup := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(grup, i)
		
	}

	grup.Wait()
	fmt.Println("Complete")
}