package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Выход по завершению выполнения функции
	go func() {
		defer func() {
			fmt.Print("First ended\n\n")
			wg.Done()
		}()
		fmt.Println("First working")

	}()
	wg.Wait()

	wg.Add(1) // Выход по таймеру через time.After
	timer := time.After(time.Second)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer:
				fmt.Print("Second ended\n\n")
				return
			default:
				fmt.Println("Second working")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	wg.Wait()

	wg.Add(1) // Выход по флагу
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Print("Third ended\n\n")
				return
			default:
				fmt.Println("Third working")
				time.Sleep(time.Millisecond)
			}
		}
	}()
	time.Sleep(time.Millisecond * 2)
	cancel()
	wg.Wait()

	wg.Add(1) // Выход по таймеру через context.WithDeadline
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Print("Fourth ended\n\n")
				return
			default:
				fmt.Println("Fourth working")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	wg.Wait()

	wg.Add(1) // Выход по таймеру через context.WithTimeout
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Print("Fifth ended\n\n")
				return
			default:
				fmt.Println("Fifth working")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	wg.Wait()

	wg.Add(1) // Выход по опустошению канала
	channel := make(chan int)
	go func() {
		defer wg.Done()
		for {
			value, ok := <-channel
			if !ok {
				fmt.Print("Sixth ended\n\n")
				return
			} else {
				fmt.Println("Sixth:", value)
			}
		}
	}()
	for i := range 5 {
		channel <- i
	}
	close(channel)
	wg.Wait()
}
