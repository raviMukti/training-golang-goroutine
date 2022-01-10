package traininggolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ { // Membuat 1000 GoRoutine
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1 // GoRoutine akan memanipulasi Data Yang Sama
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
