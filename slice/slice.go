package slice

import "errors"

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
	seen := make(map[T]struct{}, len(slice))
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}
	return result
}

// Chunk splits slice into groups of size. size must be greater than 0.
func Chunk[T any](slice []T, size int) ([][]T, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	result := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := min(i+size, len(slice))
		result = append(result, slice[i:end:end])
	}
	return result, nil
}
