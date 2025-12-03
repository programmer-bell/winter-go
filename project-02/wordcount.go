package main

import (
	"fmt"
	"strings"
	"sync"
)

func mapper(id int, text string, out chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	m := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))
	for _, w := range words {
		m[w]++
	}
	out <- m
}

func reducer(in <-chan map[string]int, done chan<- map[string]int,workers int) {
	result := make(map[string]int)
	for i := 0; i < workers; i++ {
		m := <-in
		for k, v := range m {
			result[k] += v
		}
	}
	done <- result
}

func main() {
	text := `Go is expressive, concise, clean, and efficient. Go concurrency makes it easy to build programs.`

	parts := strings.Split(text, ",") // partition
	workers := len(parts)

	mapOut := make(chan map[string]int, workers)
	var wg sync.WaitGroup

	for i, part := range parts {
		wg.Add(1)
		go mapper(i, part, mapOut, &wg)
	}
	wg.Wait()

	done := make(chan map[string]int)
	go reducer(mapOut, done, workers)

	result := <-done
	fmt.Println("Word count result:")
	fmt.Println(result)
}

