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
	if data == nil{
		return 0
	}
	count := 0
	for _, i := range data {
		count = max(i, count)
	}
	return count
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
if data == nil{
	return 0
}
	index := len(data) / CHUNKS
	remainder := len(data) % CHUNKS
	var dataPart []int
	var wg sync.WaitGroup
	wg.Add(8)
	count := 0
	for i := 0; i < 8; i++ {

		go func(count int) {
			maxInPart := 0
			for j := 0; j < index; j++ {
				maxInPart = max(data[j+count], maxInPart)
			}
			dataPart = append(dataPart, maxInPart)
			defer wg.Done()
		}(count)

		count += index
	}
	wg.Wait()
	for i := len(data) - remainder; i < len(data); i++ {
		dataPart = append(dataPart, data[i])
	}
	maximum := 0
	for _, j := range dataPart {
		maximum = max(maximum, j)
	}
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
