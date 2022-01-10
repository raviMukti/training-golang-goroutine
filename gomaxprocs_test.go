package traininggolanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()

	fmt.Println("Total CPU : ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GoRoutine : ", totalGoroutine)

	group.Wait()
}

func TestGetGomaxprocsChangeThreadNumber(t *testing.T) {

	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()

	fmt.Println("Total CPU : ", totalCpu)

	runtime.GOMAXPROCS(20) // Untuk merubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total GoRoutine : ", totalGoroutine)

	group.Wait()
}
