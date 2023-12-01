package main

//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//func countUniqueElements(arr []int, resultChan chan<- []int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// Создаем карту для отслеживания уникальных элементов
//	uniqueElements := make(map[int]struct{})
//
//	// Подсчитываем уникальные элементы в массиве
//	for _, val := range arr {
//		uniqueElements[val] = struct{}{}
//	}
//
//	// Преобразуем уникальные элементы из карты в срез
//	result := make([]int, 0, len(uniqueElements))
//	for key := range uniqueElements {
//		result = append(result, key)
//	}
//
//	// Отправляем уникальные элементы в канал
//	resultChan <- result
//}
//
//func main() {
//	// Ваш массив из 12 значений (пример случайных данных)
//	arr := generateRandomArray(12, 1)
//	arr = gaussianFilter(arr)
//
//	// Количество потоков
//	numThreads := 4
//
//	// Канал для передачи результатов от горутин
//	resultChan := make(chan []int, numThreads)
//
//	// Группа ожидания для дожидания завершения всех горутин
//	var wg sync.WaitGroup
//
//	// Разбиваем массив на части и обрабатываем каждую часть в отдельной горутине
//	for i := 0; i < numThreads; i++ {
//		wg.Add(1)
//		startIndex := i * (len(arr) / numThreads)
//		endIndex := (i + 1) * (len(arr) / numThreads)
//		go countUniqueElements(arr[startIndex:endIndex], resultChan, &wg)
//	}
//
//	// Горутина для закрытия канала после завершения всех горутин
//	go func() {
//		wg.Wait()
//		close(resultChan)
//	}()
//
//	// Карта для подсчета общего количества уникальных элементов
//	totalUniqueElements := make(map[int]struct{})
//
//	// Обработка результатов
//	for result := range resultChan {
//		for _, val := range result {
//			totalUniqueElements[val] = struct{}{}
//		}
//	}
//
//	// Выводим общее количество уникальных элементов
//	fmt.Println("Total unique elements:", len(totalUniqueElements))
//}
//
//func generateRandomArray(size, M int) []int {
//	array := make([]int, size)
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < size; i++ {
//		array[i] = rand.Intn(2*M+1) - M
//	}
//	return array
//}
