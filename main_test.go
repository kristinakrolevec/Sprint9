package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	sizes := []struct {
		lenSlice int
		want     int
	}{
		{lenSlice: 200, want: 200},
		{lenSlice: 0, want: 0},
		{lenSlice: 3000000, want: 3000000},
		{lenSlice: 1, want: 1},
	}

	for _, s := range sizes {
		slice := generateRandomElements(s.lenSlice)
		assert.Len(t, slice, s.want)
	}
}

func TestMaximum(t *testing.T) {
	slices := []struct {
		slice   []int
		wantMax int
	}{
		{slice: []int{5, 2, 8, 145, 200, 3, 1, 25}, wantMax: 200},
		{slice: nil, wantMax: 0,},
		{slice: []int{3, 3, 3}, wantMax: 3},
		{slice: []int{1}, wantMax: 1,},
	}

	for _, sl := range slices {
		max := maximum(sl.slice)
		assert.Equal(t, max, sl.wantMax)
	}
}
