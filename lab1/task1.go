package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(data []int) {
	// Подсчет в одном потоке
	startTime := time.Now()
	sequentialZeros, sequentialNegatives, sequentialPositive := countElementsSequential(data)
	sequentialTime := time.Since(startTime)

	// Подсчет в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelZeros, parallelNegatives, parallelPositives := countElementsParallel(data, numThreads)
	parallelTime := time.Since(startTime)

	// Вывод результатов и времени выполнения
	//fmt.Println("Исходный Массив:", data)
	fmt.Println("\nРезультат последовательного подсчета элементов: нули: ", sequentialZeros,
		". отрицательные эл: ", sequentialNegatives, ". Положительные эл: ", sequentialPositive)
	fmt.Println("Время выполнения последовательного подсчета элементов:", sequentialTime)

	fmt.Println("\nРезультат последовательного подсчета элементов: нули: ", parallelZeros,
		". отрицательные эл: ", parallelNegatives, ". Положительные эл: ", parallelPositives)
	fmt.Println("Время выполнения параллельного подсчета элементов:", parallelTime)
}

func countElementsSequential(data []int) (int, int, int) {
	var zeros, negatives, positives int
	for _, value := range data {
		switch {
		case value == 0:
			zeros++
		case value < 0:
			negatives++
		case value > 0:
			positives++
		}
	}
	return zeros, negatives, positives
}

func countElementsParallel(data []int, numThreads int) (int, int, int) {
	var zeros, negatives, positives int
	chunkSize := len(data) / numThreads

	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()
			localZeros, localNegatives, localPositives := 0, 0, 0
			for j := start; j < end; j++ {
				switch {
				case data[j] == 0:
					localZeros++
				case data[j] < 0:
					localNegatives++
				case data[j] > 0:
					localPositives++
				}
			}

			// Используем мьютекс для безопасного доступа к общим результатам
			var mutex sync.Mutex
			mutex.Lock()
			zeros += localZeros
			negatives += localNegatives
			positives += localPositives
			mutex.Unlock()
		}(start, end)
	}

	wg.Wait()
	return zeros, negatives, positives
}
