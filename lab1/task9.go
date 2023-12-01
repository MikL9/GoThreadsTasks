package main

import (
	"fmt"
	"sync"
	"time"
)

func task9() {
	matrixA := [][]float64{
		{10, -5, 2, 3, 8},
		{-2, 15, 4, -7, 6},
		{3, 1, 20, -4, 7},
		{6, -9, 8, 25, -12},
		{-4, 7, 5, 6, 18},
	}

	vectorB := []float64{23, -8, 15, 32, -10}

	// Неправильное (с гонками) решение системы
	startTime := time.Now()
	incorrectResult := solveLinearSystemIncorrect(matrixA, vectorB)
	parallelTime := time.Since(startTime)
	fmt.Println("Время выполнения с гонками: ", parallelTime)
	fmt.Println("Неправильное решение системы:")
	fmt.Println(incorrectResult)

	// Правильное решение системы (с синхронизацией)
	startTime = time.Now()
	correctResult := solveLinearSystemCorrect(matrixA, vectorB)
	parallelTime = time.Since(startTime)
	fmt.Println("\nВремя выполнения без гонок: ", parallelTime)
	fmt.Println("Правильное решение системы:")
	fmt.Println(correctResult)
}

func generateRandomMatrixV2(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
		for j := range matrix[i] {
			matrix[i][j] = float64(i*cols+j+1) + float64(j)/1000.0
		}
	}
	return matrix
}

// Генерация случайного вектора
func generateRandomVector(size int) []float64 {
	vector := make([]float64, size)
	for i := range vector {
		vector[i] = float64(i) + float64(i)/1000.0
	}
	return vector
}

// Неправильное (с гонками) решение системы линейных уравнений
func solveLinearSystemIncorrect(matrixA [][]float64, vectorB []float64) []float64 {
	rows := len(matrixA)
	cols := len(matrixA[0])

	result := make([]float64, cols)

	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := i + 1; j < rows; j++ {
				ratio := matrixA[j][i] / matrixA[i][i]
				for k := i; k < cols; k++ {
					matrixA[j][k] -= ratio * matrixA[i][k]
				}
				vectorB[j] -= ratio * vectorB[i]
			}
		}(i)
	}

	wg.Wait()

	for i := rows - 1; i >= 0; i-- {
		result[i] = vectorB[i] / matrixA[i][i]
		for j := i + 1; j < cols; j++ {
			result[i] -= matrixA[i][j] / matrixA[i][i] * result[j]
		}
	}

	return result
}

// Правильное решение системы линейных уравнений (с синхронизацией)
func solveLinearSystemCorrect(matrixA [][]float64, vectorB []float64) []float64 {
	rows := len(matrixA)
	cols := len(matrixA[0])

	result := make([]float64, cols)

	var mutex sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := i + 1; j < rows; j++ {
				ratio := matrixA[j][i] / matrixA[i][i]
				for k := i; k < cols; k++ {
					matrixA[j][k] -= ratio * matrixA[i][k]
				}
				vectorB[j] -= ratio * vectorB[i]
			}

			// Синхронизированный доступ к результату
			mutex.Lock()
			for i := rows - 1; i >= 0; i-- {
				result[i] = vectorB[i] / matrixA[i][i]
				for j := i + 1; j < cols; j++ {
					result[i] -= matrixA[i][j] / matrixA[i][i] * result[j]
				}
			}
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	return result
}
