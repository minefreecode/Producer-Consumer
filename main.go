package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Паттерн producer-consumer ===")

	bufferSize := 5
	numProducers := 2
	numConsumers := 3
	numItems := 10

	buffer := make(chan int, bufferSize)
	var wg sync.WaitGroup

	for p := 1; p <= numProducers; p++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < numItems; i++ {
				item := rand.Intn(100)
				buffer <- item
				fmt.Printf("Продюсер %d произвёл: %d\n", id, item)
				time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
			}
		}(p)
	}

	var consumerWg sync.WaitGroup
	for c := 1; c <= numConsumers; c++ {
		consumerWg.Add(1) // Увеличивает на 1
		go func(id int) {
			defer consumerWg.Done() //Умиеньшает на 1
			for item := range buffer {
				fmt.Printf("Потребитель %d потребил: %d\n", id, item)
				time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
			}
		}(c)
	}

	wg.Wait()
	close(buffer)

	consumerWg.Wait()

	fmt.Println("Продюсер-Консумер завершены")
}
