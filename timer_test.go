package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T)  {
	timer := time.NewTimer(5 * time.Second)	
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T)  {
	channel := time.After(5 * time.Second)	
	fmt.Println(time.Now())

	fmt.Println(<- channel)
}

func TestAfterFunc(t *testing.T)  {
	wg := sync.WaitGroup {}
	wg.Add(1)

	fmt.Println(time.Now())
	time.AfterFunc(3 * time.Second, func() {
		defer wg.Done()
		fmt.Println("Execute after 3 second")
		fmt.Println(time.Now())
	})

	wg.Wait()
}