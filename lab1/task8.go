package main

import (
	"fmt"
	"sync"
	"time"
)

func task8() {
	// Задаем размеры матриц
	rowsA, colsA := 510, 510
	rowsB, colsB := 510, 510

	// Инициализируем матрицы случайными значениями
	matrixA := generateRandomMatrix(rowsA, colsA)
	matrixB := generateRandomMatrix(rowsB, colsB)

	//без потоков
	startTime := time.Now()
	_ = multiplyMatrixSingleThread(matrixA, matrixB)
	parallelTime := time.Since(startTime)
	fmt.Println("Умножение матриц без потоков:")
	fmt.Println("Время выполнения:", parallelTime)

	// Неправильное (с гонками) умножение матриц
	startTime = time.Now()
	_ = multiplyMatrixIncorrect(matrixA, matrixB)
	parallelTime = time.Since(startTime)
	fmt.Println("\nНеправильное умножение матриц:")
	fmt.Println("Время выполнения с гонками:", parallelTime)
	//printMatrix(incorrectResult)

	// Правильное умножение матриц (с синхронизацией)
	startTime = time.Now()
	_ = multiplyMatrixCorrect(matrixA, matrixB)
	parallelTime = time.Since(startTime)
	fmt.Println("\nПравильное умножение матриц:")
	fmt.Println("Время выполнения без гонок:", parallelTime)
	//printMatrix(correctResult)
}

func generateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = i*cols + j + 1
		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Неправильное умножение матриц (с гонками)
func multiplyMatrixIncorrect(matrixA, matrixB [][]int) [][]int {
	rowsA, colsA := len(matrixA), len(matrixA[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
	}

	var wg sync.WaitGroup

	for i := 0; i < rowsA; i++ {
		for j := 0; j < len(matrixB[0]); j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				for k := 0; k < colsA; k++ {
					result[i][j] += matrixA[i][k] * matrixB[k][j]
				}
			}(i, j)
		}
	}

	wg.Wait()
	return result
}

// Правильное умножение матриц (с синхронизацией)
func multiplyMatrixCorrect(matrixA, matrixB [][]int) [][]int {
	rowsA, colsA := len(matrixA), len(matrixA[0])
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, len(matrixB[0]))
	}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < rowsA; i++ {
		for j := 0; j < len(matrixB[0]); j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				localResult := 0
				for k := 0; k < colsA; k++ {
					localResult += matrixA[i][k] * matrixB[k][j]
				}

				// Синхронизированный доступ к результату
				mutex.Lock()
				result[i][j] = localResult
				mutex.Unlock()
			}(i, j)
		}
	}
	wg.Wait()
	return result
}

func multiplyMatrixSingleThread(matrixA, matrixB [][]int) [][]int {
	rowsA, colsA := len(matrixA), len(matrixA[0])
	rowsB, colsB := len(matrixB), len(matrixB[0])

	if colsA != rowsB {
		panic("Невозможно умножить матрицы: количество столбцов первой матрицы не равно количеству строк второй матрицы")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}
