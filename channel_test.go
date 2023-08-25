package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Abrar Naim"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Abrar Naim"
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	// close(channel)
}

func OnlyIn(channel chan <- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Abrar Naim"
}

func OnlyOut(channel <- chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	// close(channel)
}

func TestBufferdChannel(t *testing.T)  {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Abrar"
		channel <- "Naim"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T)  {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data" + data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t * testing.T)  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("Data Dari chanel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data Dari chanel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t * testing.T)  {
	t.Skip("skip kebanyakn Println")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <- channel1:
			fmt.Println("Data Dari chanel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data Dari chanel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}