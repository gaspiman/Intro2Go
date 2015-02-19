package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go race(i, m, wg)
	}
	wg.Wait()
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func race(i int, m map[int]int, wg *sync.WaitGroup) {
	m[i] = i
	wg.Done()
}
