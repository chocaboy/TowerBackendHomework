package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // Чтобы программа выполнялась конкурентно (по заданию), для параллельного выполнения эту строчку нужно удалить
	var wg sync.WaitGroup
	a := []int{2, 4, 6, 8, 10}
	for _, temp := range a {
		wg.Add(1)
		go func(temp int) {
			defer wg.Done()
			fmt.Println(temp * temp)
		}(temp)
	}
	wg.Wait()
}
