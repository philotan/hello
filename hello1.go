package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go func() {
		fmt.Println("Goroutine отправляет 10, 20, 30")
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
		fmt.Println("Goroutine закрыла канал.")
	}()

	fmt.Println("Main принимает данные...")
	for value := range ch {
		fmt.Printf("Main получил: %d\n", value)
	}
	fmt.Println("Main завершил прием.")
}
