package arraysAndSlice

import (
	"reflect"
	"strings"
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

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		assertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		assertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya         = Account{Name: "Riya", Balance: 100}
		chris        = Account{Name: "Chris", Balance: 75}
		adil         = Account{Name: "Adil", Balance: 200}
		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)
	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(transactions, account).Balance
	}
	assertEqual(t, newBalanceFor(riya), 200)
	assertEqual(t, newBalanceFor(chris), 0)
	assertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		assertEqual(t, found, true)
		assertEqual(t, firstEvenNumber, 2)
	})
	type Person struct {
		Name string
	}
	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			Person{Name: "Kent Beck"},
			Person{Name: "Martin Fowler"},
			Person{Name: "Chris James"},
		}
		king, found := Find(people, func(x Person) bool {
			return strings.Contains(x.Name, "Chris")
		})
		assertEqual(t, found, true)
		assertEqual(t, king, Person{Name: "Chris James"})
	})
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
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
