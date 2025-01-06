package arraysAndSlice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		assertResult(t, got, want, nums)
	})
}

func TestSumAll(t *testing.T) {
	input := [][]int{{1, 2}, {0, 9}}
	got := SumAll(input[0], input[1])
	want := []int{3, 9}
	assertSliceResult(t, got, want, input[:])
}

func TestSumAllTails(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		input := [][]int{{1, 2}, {0, 9}}
		got := SumAllTails(input[0], input[1])
		want := []int{2, 9}
		assertSliceResult(t, got, want, input[:])
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		input := [][]int{{}, {3, 4, 5}}
		got := SumAllTails(input...)
		want := []int{0, 9}
		assertSliceResult(t, got, want, input[:])
	})
}

func assertResult(t *testing.T, got, want int, given []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, given)
	}
}

func assertSliceResult(t *testing.T, got, want []int, given [][]int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, given)
	}
}
