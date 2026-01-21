package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	channel := make(chan int)
	print("Введите число секунд выделенное для программы: ")
	fmt.Scan(&n)
	timer := time.After(time.Duration(n) * time.Second) // Таймер на n секунд
	go func() {                                         // Запись в канал
		defer close(channel)
		temp := 1
		for {
			select {
			case <-timer:
				fmt.Println("Время вышло")
				return
			case channel <- temp:
				temp++
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()
	for temp := range channel { // Вывод из канала
		fmt.Println("Число из канала:", temp)
	}
}
