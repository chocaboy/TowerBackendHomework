package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var n int
	var wg sync.WaitGroup
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	channel := make(chan string)
	fmt.Print("Count of workers: ")
	fmt.Scan(&n)
	wg.Add(1)

	go func() { // Ввод
		var temp string
		for {
			fmt.Print("Enter string: ")
			fmt.Scan(&temp)
			channel <- temp
		}
	}()

	for i := 0; i < n; i++ { // Запуск n воркеров вывода
		go func() {
			defer wg.Done()
			for {
				select {
				case out := <-channel:
					fmt.Println("you entered:", out)
				case err := <-sig: // Проверка на Ctrl + C
					fmt.Println("\nProgram ended with:", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
