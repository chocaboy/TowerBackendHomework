package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(1) // Чтобы программа выполнялась конкурентно (по заданию), для параллельного выполнения эту строчку нужно удалить
	go_map := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			go_map[i] = i
		}()
	}
	wg.Wait()
	fmt.Println(go_map)
}
