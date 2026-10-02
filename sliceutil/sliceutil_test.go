package sliceutil

import (
	"reflect"
	"strconv"
	"testing"
)

func TestContains(t *testing.T) {
	strings := []string{"a", "b", "c"}
	if !Contains(strings, "b") {
		t.Errorf("Contains() should return true for existing element")
	}
	if Contains(strings, "z") {
		t.Errorf("Contains() should return false for non-existing element")
	}

	ints := []int{1, 2, 3}
	if !Contains(ints, 2) {
		t.Errorf("Contains() should return true for existing element")
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 1})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unique() = %v; want %v", got, want)
	}
}

func TestChunk(t *testing.T) {
	got, err := Chunk([]int{1, 2, 3, 4, 5}, 2)
	if err != nil {
		t.Fatalf("Chunk() unexpected error: %v", err)
	}
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunk() = %v; want %v", got, want)
	}

	if _, err := Chunk([]int{1}, 0); err == nil {
		t.Error("Chunk() with size 0 expected error")
	}
}

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3}
	strs := Map(ints, func(i int) string {
		return strconv.Itoa(i)
	})

	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(strs, expected) {
		t.Errorf("Map() = %v; want %v", strs, expected)
	}
}
