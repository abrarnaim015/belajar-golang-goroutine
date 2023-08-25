package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// var locker = sync.Mutex {}
// var cond = sync.NewCond(&locker)
// var group = sync.WaitGroup {}

// func WaitCondition(value int)  {
// 	defer group.Done()
// 	group.Add(1)

// 	cond.L.Lock()
// 	cond.Wait()
// 	fmt.Println("Done", value)
// 	cond.L.Unlock()
// }

func TestCond(t *testing.T)  {
	cond := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup {}
	
	WaitCondition := func (value int)  {
		defer wg.Done()
		cond.L.Lock()
		cond.Wait()
		fmt.Println("Done", value)
		cond.L.Unlock()
	}


	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCondition(i)
	}
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // untuk nge unlock "cond.Wait()" per sekali panggil
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() // untuk nge unlock "cond.Wait()" seterusnya, alhasin "cond.Wait()" tidak berfungsi lg
	// }()

	wg.Wait()
}