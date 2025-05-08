package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size == 0 {
		return nil
	}
	slice := make([]int, size)
	scr := rand.NewSource(time.Now().Unix())
	for i := 0; i < size; i++ {
		slice[i] = int(scr.Int63())
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if data == nil {
		return 0
	}
	maximum := 0
	for _, num := range data {
		maximum = max(num, maximum)
	}
	return maximum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if data == nil {
		fmt.Println("Slice is empty")
		return 0
	}
	index := len(data) / CHUNKS

	var dataPart []int
	var newData []int
	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	count := 0
	for i := 0; i < CHUNKS; i++ {
		go func(count int) {
			defer wg.Done()
			if i == 7 {
				dataPart = data[count:]
			} else {
				dataPart = data[count:(count + index)]
			}
			maxInPart := maximum(dataPart)
			newData = append(newData, maxInPart)
		}(count)

		count += index

	}
	wg.Wait()
	maximum := maximum(newData)
	return maximum
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	numbersSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	now := time.Now()
	max := maximum(numbersSlice)
	elapsed := time.Now().Sub(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	now = time.Now()
	max = maxChunks(numbersSlice)
	elapsed = time.Now().Sub(now)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
