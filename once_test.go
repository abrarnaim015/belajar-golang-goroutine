package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce()  {
	counter++
}

func TestOnce(t *testing.T)  {
	once := sync.Once {}
	group := sync.WaitGroup {}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) // untuk memanggil func hanya 1 kali, jika sudah di panggil next tidak akan ke panggil lagi
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", counter)
}