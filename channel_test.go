package traininggolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// channel <- "Ravi" // Kirim value ke channel

	// data := <-channel // Ambil value dari channel

	// fmt.Println(<-channel) // Dikirim langsung ke parameter

	// defer close(channel) // untuk menutup channel

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Ravi Mukti"

		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ravi Mukti"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Ravi Mukti"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Ravi"
		channel <- "Mukti"
		channel <- "Hartadi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(channel) // harus tetap close channel
	}()

	for data := range channel {
		fmt.Println("Menerima Data ", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select {
	// case data := <-channel1:
	// 	fmt.Println("Data Dari Channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("Data Dari Channel 2", data)
	// }

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
