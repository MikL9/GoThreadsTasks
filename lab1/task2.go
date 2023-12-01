package main

import (
	"fmt"
	"sync"
	"time"
)

func task2(data []int) {
	// Подсчет в одном потоке
	startTime := time.Now()
	sequentialMean := meanSequential(data)
	sequentialTime := time.Since(startTime)

	// Подсчет в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelMean := meanParallel(data, numThreads)
	parallelTime := time.Since(startTime)

	// Вывод результатов и времени выполнения
	fmt.Println("\nРезультат последовательного вычисления среднего значения: ", sequentialMean)
	fmt.Println("Время выполнения последовательного вычисления: ", sequentialTime)

	fmt.Println("\nРезультат параллельного вычисления среднего значения: ", parallelMean)
	fmt.Println("Время выполнения параллельного вычисления: ", parallelTime)
}

func meanSequential(data []int) float64 {
	var sum float64

	for _, value := range data {
		sum += float64(value)
	}

	mean := sum / float64(len(data))
	return mean
}

func meanParallel(data []int, numThreads int) float64 {
	var sum float64
	chunkSize := len(data) / numThreads

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()
			localSum := 0.0

			for j := start; j < end; j++ {
				localSum += float64(data[j])
			}

			// Используем мьютекс для безопасного доступа к общим результатам
			var mutex sync.Mutex
			mutex.Lock()
			sum += localSum
			mutex.Unlock()
		}(start, end)
	}

	wg.Wait()

	mean := sum / float64(len(data))
	return mean
}
