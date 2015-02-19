package main

import (
	"fmt"
	"sync"
)

type info struct {
	key   int
	value int
}

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan *info)
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go race(i, wg, ch)
	}
	go func() {
		for g := range ch {
			m[g.key] = g.value
		}
	}()
	wg.Wait()
	close(ch)
	fmt.Println("DONE! - ALL Clear")
}

func race(i int, wg *sync.WaitGroup, ch chan *info) {
	g := info{key: i, value: i}
	ch <- &g
	wg.Done()
}
