package main

import (
	"fmt"
	"sync"
	"time"
)

func task7(data []int) {
	data = removeNegatives(data)
	// Поиск в одном потоке
	startTime := time.Now()
	sortSequence := countingSort(data)
	sequentialTime := time.Since(startTime)

	fmt.Println("\nСортировка (последовательно): ", sortSequence[:30])
	fmt.Println("Время выполнения поиска (последовательно): ", sequentialTime)

	// Поиск в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelSort := countingSortParallel(data, numThreads)
	parallelTime := time.Since(startTime)

	fmt.Println("\nСортировка (параллельно): ", parallelSort[:30])
	fmt.Println("Время выполнения поиска (параллельно): ", parallelTime)
}

func removeNegatives(arr []int) []int {
	var result []int
	for _, val := range arr {
		if val >= 0 {
			result = append(result, val)
		}
	}
	return result
}

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Находим минимум и максимум в массиве
	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	// Создаем массив для подсчета частоты встречаемости каждого элемента
	count := make([]int, max-min+1)

	// Подсчитываем частоту встречаемости каждого элемента
	for _, val := range arr {
		count[val-min]++
	}

	// Строим отсортированный массив на основе частоты встречаемости
	sortedArr := make([]int, 0, len(arr))
	for i, freq := range count {
		for j := 0; j < freq; j++ {
			sortedArr = append(sortedArr, i+min)
		}
	}

	return sortedArr
}

func countingSortParallel(arr []int, numThreads int) []int {
	if len(arr) == 0 {
		return arr
	}

	min, max := arr[0], arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	count := make([]int, max-min+1)

	chunkSize := len(arr) / numThreads
	var wg sync.WaitGroup
	wg.Add(numThreads)

	for i := 0; i < numThreads; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()
			localCount := make([]int, max-min+1)

			for j := start; j < end; j++ {
				localCount[arr[j]-min]++
			}

			for j := range localCount {
				count[j] += localCount[j]
			}
		}(start, end)
	}
	wg.Wait()

	sortedArr := make([]int, 0, len(arr))
	for i, freq := range count {
		for j := 0; j < freq; j++ {
			sortedArr = append(sortedArr, i+min)
		}
	}

	return sortedArr
}
