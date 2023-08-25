package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld()  {
	fmt.Println("Hello World")
}

func TestCreateForoutine(t *testing.T)  {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int)  {
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T)  {
	for i := 0; i < 10; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}