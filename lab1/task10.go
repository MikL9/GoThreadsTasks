package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task10() {
	dataSize := 1000000
	data := generateRandomArrayFloat(dataSize)

	// Задаем случайное число для умножения
	rand.Seed(time.Now().UnixNano())
	multiplier := rand.Float64()

	// Обработка массива с гонками
	startTimeRace := time.Now()
	resultRace := processArrayWithRace(data, multiplier)
	raceTime := time.Since(startTimeRace)

	// Обработка массива без гонок
	startTimeNoRace := time.Now()
	resultNoRace := processArrayWithoutRace(data, multiplier)
	noRaceTime := time.Since(startTimeNoRace)

	// Вывод результатов
	fmt.Println("Гонки (Race):", resultRace[:50])
	fmt.Println("Время выполнения (Race):", raceTime)

	fmt.Println("\nБез гонок (No Race):", resultNoRace[:50])
	fmt.Println("Время выполнения (No Race):", noRaceTime)
}

func generateRandomArrayFloat(size int) []float64 {
	arr := make([]float64, size)
	for i := range arr {
		arr[i] = float64(i + 1)
	}
	return arr
}

// Обработка массива с гонками
func processArrayWithRace(data []float64, multiplier float64) []float64 {
	var result []float64
	var wg sync.WaitGroup

	for _, value := range data {
		wg.Add(1)
		go func(value float64) {
			defer wg.Done()
			// Гонка! Нет синхронизации при доступе к общим данным
			result = append(result, value*multiplier)
		}(value)
	}

	wg.Wait()
	return result
}

// Обработка массива без гонок
func processArrayWithoutRace(data []float64, multiplier float64) []float64 {
	var result []float64
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, value := range data {
		wg.Add(1)
		go func(value float64) {
			defer wg.Done()
			// Без гонок! Используем мьютекс для синхронизации
			mutex.Lock()
			result = append(result, value*multiplier)
			mutex.Unlock()
		}(value)
	}

	wg.Wait()
	return result
}
