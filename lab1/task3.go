package main

import (
	"fmt"
	"sync"
	"time"
)

func task3(data []int) {
	// Подсчет в одном потоке
	startTime := time.Now()
	sequentialChains := countChainsSequential(data, 10)
	sequentialTime := time.Since(startTime)

	// Подсчет в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelChains := countChainsParallel(data, 10, numThreads)
	parallelTime := time.Since(startTime)

	// Вывод результатов и времени выполнения
	fmt.Println("\nРезультат последовательного подсчета цепочек из 10 нулей: ", sequentialChains)
	fmt.Println("Время выполнения последовательного подсчета: ", sequentialTime)

	fmt.Println("\nРезультат параллельного подсчета цепочек из 10 нулей: ", parallelChains)
	fmt.Println("Время выполнения параллельного подсчета: ", parallelTime)
}

func countChainsSequential(data []int, k int) int {
	var chains int
	var consecutiveZeros int

	for _, value := range data {
		if value == 0 {
			consecutiveZeros++
			if consecutiveZeros == k {
				chains++
			}
		} else {
			consecutiveZeros = 0
		}
	}

	return chains
}

func countChainsParallel(data []int, k, numThreads int) int {
	var chains int
	chunkSize := len(data) / numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()

			localChains := 0
			localConsecutiveZeros := 0

			for j := start; j < end; j++ {
				if data[j] == 0 {
					localConsecutiveZeros++
					if localConsecutiveZeros == k {
						localChains++
					}
				} else {
					localConsecutiveZeros = 0
				}
			}

			var mutex sync.Mutex
			mutex.Lock()
			chains += localChains
			mutex.Unlock()
		}(start, end)
	}

	wg.Wait()
	return chains
}
