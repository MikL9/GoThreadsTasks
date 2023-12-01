package main

import (
	"fmt"
	"sync"
	"time"
)

func task4(data []int) {
	// Поиск в одном потоке
	startTime := time.Now()
	_ = findLongestSequence(data)
	sequentialTime := time.Since(startTime)

	// Поиск в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelSequence := findLongestSequenceParallel(data, numThreads)
	parallelTime := time.Since(startTime)

	// Вывод результатов и времени выполнения
	//fmt.Println("\nСамая длинная последовательность элементов с суммой 0 (последовательно): ", sequentialSequence)
	fmt.Println("Время выполнения поиска (последовательно): ", sequentialTime)

	fmt.Println("\nСамая длинная последовательность элементов с суммой 0 (параллельно): ", len(parallelSequence))
	fmt.Println("Время выполнения поиска (параллельно): ", parallelTime)
}

func findLongestSequence(data []int) []int {
	var longestSequence []int
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			subsequence := data[i : j+1]
			if sum(subsequence) == 0 && len(subsequence) > len(longestSequence) {
				longestSequence = subsequence
			}
		}
	}
	return longestSequence
}

func findLongestSequenceParallel(data []int, numThreads int) []int {
	var longestSequence []int
	chunkSize := len(data) / numThreads

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()

			localLongestSequence := findLongestSequence(data[start:end])

			var mutex sync.Mutex
			mutex.Lock()
			if len(localLongestSequence) > len(longestSequence) {
				longestSequence = localLongestSequence
			}
			mutex.Unlock()
		}(start, end)
	}

	wg.Wait()

	return longestSequence
}

func sum(slice []int) int {
	result := 0
	for _, value := range slice {
		result += value
	}
	return result
}
