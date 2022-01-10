package traininggolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ { // Membuat 1000 GoRoutine
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()   // Akan lock goroutine
				x = x + 1      // GoRoutine akan memanipulasi Data Yang Sama
				mutex.Unlock() // melepaskan goroutine
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
