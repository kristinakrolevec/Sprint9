package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	for i := 0; i < 100; i++ {
		slice := generateRandomElements(i)
		assert.Len(t, slice, i)
	}
}

func TestMaximum(t *testing.T) {
	slice := make([]int, 200)
	for i := 0; i < 200; i++ {
		slice[i] = i
	}
	max := maximum(slice)
	assert.Equal(t, max, 199)

	slice = nil
	max = maximum(slice)
	assert.Equal(t, max, 0)
}


