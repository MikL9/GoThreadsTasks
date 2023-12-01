package main

import (
	"fmt"
	"sync"
	"time"
)

func task5(data []int) {
	// Подсчет в одном потоке
	startTime := time.Now()
	sequentialCount := countUniqueSequential(data)
	sequentialTime := time.Since(startTime)

	// Вывод результатов и времени выполнения
	fmt.Println("\nКоличество уникальных элементов (последовательно): ", sequentialCount)
	fmt.Println("Время выполнения подсчета (последовательно): ", sequentialTime)

	// Подсчет в нескольких потоках
	startTime = time.Now()
	numThreads := 4
	parallelCount := countUniqueParallel(data, numThreads)
	parallelTime := time.Since(startTime)

	fmt.Println("\nКоличество уникальных элементов (параллельно): ", parallelCount)
	fmt.Println("Время выполнения подсчета (параллельно): ", parallelTime)
}

func countUniqueSequential(data []int) int {
	count := 0
	occurrences := make(map[int]int)

	for _, value := range data {
		occurrences[value]++
	}

	for _, occurrence := range occurrences {
		if occurrence == 1 {
			count++
		}
	}

	return count
}

func findUniqueElements(arr []int, resultChan chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	frequencyMap := make(map[int]int)

	for _, val := range arr {
		frequencyMap[val]++
	}

	uniqueElements := make([]int, 0)

	for key, count := range frequencyMap {
		if count == 1 {
			uniqueElements = append(uniqueElements, key)
		}
	}

	resultChan <- uniqueElements
}

func countUniqueParallel(data []int, numThreads int) int {
	arr := data
	resultChan := make(chan []int, numThreads)

	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		startIndex := i * (len(arr) / numThreads)
		endIndex := (i + 1) * (len(arr) / numThreads)
		go findUniqueElements(arr[startIndex:endIndex], resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	uniqueElementsMap := make(map[int]struct{})
	for result := range resultChan {
		for _, val := range result {
			uniqueElementsMap[val] = struct{}{}
		}
	}
	uniqueElements := make([]int, 0, len(uniqueElementsMap))
	for val := range uniqueElementsMap {
		uniqueElements = append(uniqueElements, val)
	}

	dataOccurrences := make(map[int]int)
	for _, val := range data {
		dataOccurrences[val]++
	}
	var uniqueElementsOccurringOnce []int
	for _, val := range uniqueElements {
		if count, found := dataOccurrences[val]; found && count == 1 {
			uniqueElementsOccurringOnce = append(uniqueElementsOccurringOnce, val)
		}
	}

	return len(uniqueElementsOccurringOnce)
}
