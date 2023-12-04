package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	arraySize := 100000000
	data := generateRandomArray(arraySize)

	randomMultiplier := rand.Intn(5) + 1

	// Используем общую переменную для агрегации результатов
	var resultCommon int
	startTimeCommon := time.Now()
	processArrayCommon(data, &resultCommon, randomMultiplier)
	commonTime := time.Since(startTimeCommon)
	fmt.Println("Метод 1 (с использованием общей переменной):")
	fmt.Println("Результат:", resultCommon)
	fmt.Println("Время выполнения:", commonTime)

	// Используем атомарную переменную для агрегации результатов
	var resultAtomic int64
	startTimeAtomic := time.Now()
	processArrayAtomic(data, &resultAtomic, randomMultiplier)
	atomicTime := time.Since(startTimeAtomic)
	fmt.Println("\nМетод 2 (с использованием атомарных операций):")
	fmt.Println("Результат:", atomic.LoadInt64(&resultAtomic))
	fmt.Println("Время выполнения:", atomicTime)

	// Используем индивидуальные переменные для каждой горутины
	var resultIndividuals []int
	startTimeIndividual := time.Now()
	processArrayIndividual(data, &resultIndividuals, randomMultiplier)
	individualTime := time.Since(startTimeIndividual)
	fmt.Println("\nМетод 4 (с использованием индивидуальных переменных для каждого потока):")
	fmt.Println("Результат:", resultIndividuals[:10])
	fmt.Println("Время выполнения:", individualTime)

	// Используем мьютекс для многопоточного метода
	var resultMutexConcurrent int
	var mutexConcurrent sync.Mutex
	startTimeMutexConcurrent := time.Now()
	numThreads := 4
	processArrayMutexConcurrent(data, &resultMutexConcurrent, &mutexConcurrent, randomMultiplier, numThreads)
	mutexConcurrentTime := time.Since(startTimeMutexConcurrent)
	fmt.Println("\nМетод 5 (многопоточный метод с использованием мьютекса):")
	fmt.Println("Результат:", resultMutexConcurrent)
	fmt.Println("Время выполнения:", mutexConcurrentTime)
}

// Генерация случайного массива
func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, size)
	for i := range array {
		array[i] = rand.Intn(100)
	}
	return array
}

// Умножение каждого элемента массива на случайное число с использованием общей переменной
func processArrayCommon(data []int, result *int, randomMultiplier int) {
	for _, value := range data {
		*result += value * randomMultiplier
	}
}

// Умножение каждого элемента массива на случайное число с использованием атомарных операций
func processArrayAtomic(data []int, result *int64, randomMultiplier int) {
	for _, value := range data {
		atomic.AddInt64((*int64)(unsafe.Pointer(result)), int64(value*randomMultiplier))
	}
}

// Умножение каждого элемента массива на случайное число с использованием индивидуальных переменных для каждого потока
func processArrayIndividual(data []int, result *[]int, randomMultiplier int) {
	localResults := make([]int, len(data))
	for i, value := range data {
		localResults[i] = value * randomMultiplier
	}
	*result = localResults
}

func processArrayMutexConcurrent(data []int, result *int, mutex *sync.Mutex, randomMultiplier int, numThreads int) {
	var wg sync.WaitGroup
	wg.Add(numThreads)

	chunkSize := len(data) / numThreads

	for i := 0; i < numThreads; i++ {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize

		go func(start, end int) {
			defer wg.Done()
			localResult := 0
			for j := start; j < end; j++ {
				localResult += data[j] * randomMultiplier
			}

			mutex.Lock()
			*result += localResult
			mutex.Unlock()
		}(startIndex, endIndex)
	}

	wg.Wait()
}
