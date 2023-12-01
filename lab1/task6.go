package main

import "fmt"

func task6(data []int) {
	smoothedData := gaussianFilter(data)
	fmt.Println("Original Data:", data[:10])
	fmt.Println("Smoothed Data:", smoothedData[:10])
}

func gaussianFilter(data []int) []int {
	size := len(data)
	result := make([]int, size)

	kernel := []float64{1, 2, 1, 2, 4, 2, 1, 2, 1}

	sum := 0.0
	for _, value := range kernel {
		sum += value
	}
	for i := range kernel {
		kernel[i] /= sum
	}

	for i := 0; i < size; i++ {
		for j := -1; j <= 1; j++ {
			index := i + j
			if index >= 0 && index < size {
				result[i] += int(float64(data[index]) * kernel[j+1])
			}
		}
	}

	return result
}
