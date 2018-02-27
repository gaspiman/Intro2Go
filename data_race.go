package main

import (
	"fmt"
	"sync"
)

func race_main() {
	m := make(map[int]int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go race_main_worker(i, m, wg)
	}
	wg.Wait()
	fmt.Println("DONE! - ALL Clear")
}

func race_main_worker(i int, m map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	m[i] = i
}
