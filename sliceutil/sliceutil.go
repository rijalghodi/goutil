package sliceutil

import (
	"github.com/rijalghodi/goutil/internal/validate"
	"github.com/samber/lo"
)

// Contains returns true if the given slice contains the target element.
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// Unique returns a new slice with duplicate values removed, keeping first occurrence order.
func Unique[T comparable](slice []T) []T {
	return lo.Uniq(slice)
}

// Chunk splits slice into groups of size. size must be greater than 0.
func Chunk[T any](slice []T, size int) ([][]T, error) {
	if err := validate.Positive(size, "size"); err != nil {
		return nil, err
	}
	return lo.Chunk(slice, size), nil
}
