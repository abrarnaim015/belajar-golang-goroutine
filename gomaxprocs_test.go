package belajargolanggoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T)  {
	totalCpu := runtime.NumCPU()
	totalThread := runtime.GOMAXPROCS(-1)
	totalGoroutine := runtime.NumGoroutine()

	fmt.Println("totalCpu :", totalCpu, "\n", "totalThread :", totalThread, "\n", "totalGoroutine :", totalGoroutine)
}