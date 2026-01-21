package main

import (
	"fmt"
	"os"
)

func masToChan(a []int) chan int { // Функция переноса чисел из массива в канал
	channel := make(chan int)
	go func() {
		defer close(channel)
		for i := range a {
			channel <- i
		}
	}()
	return channel
}

func chanToSq(a <-chan int) chan int { // Функция переноса чисел из канала в другой канал с возведением в квадрат
	channel := make(chan int)
	go func() {
		defer close(channel)
		for i := range a {
			channel <- i * i
		}
	}()
	return channel
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	output := chanToSq(masToChan(arr))
	for i := range output {
		fmt.Fprintln(os.Stdout, i) //тоже самое что и fmt.Println()
	}
}
