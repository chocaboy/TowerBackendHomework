package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // Чтобы программа выполнялась конкурентно (по заданию), для параллельного выполнения эту строчку нужно удалить
	var wg sync.WaitGroup
	var mu sync.Mutex
	a := []int{2, 4, 6, 8, 10}
	var sum int = 0
	for _, temp := range a {
		wg.Add(1)
		go func(int) {
			defer wg.Done()
			mu.Lock()
			sum += temp * temp
			mu.Unlock()
		}(temp)
	}
	wg.Wait()
	fmt.Println(sum)
}
