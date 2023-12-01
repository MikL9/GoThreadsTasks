package main

import (
	"math/rand"
	"time"
)

func main() {
	//arraySize := 100000000
	arraySize := 100000000
	M := 100

	_ = generateRandomArray(arraySize, M)

	//task1(data)
	//task2(data)
	//task3(data)
	//task4(data)
	//task5(data)
	//task6(data)
	//task7(data)
	//task8()
	//task9()
	task10()
}

func generateRandomArray(size, M int) []int {
	array := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(2*M+1) - M
	}
	return array
}
