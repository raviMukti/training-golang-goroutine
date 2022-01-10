package traininggolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Ravi")
	pool.Put("Mukti")
	pool.Put("Hartadi")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // ALways put again the data, neither it will remove from pool
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
