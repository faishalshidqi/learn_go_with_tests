package arraysAndSlice

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Sum(numbers []int) int {
	/*
		sum := 0
		for _, i := range numbers {
			sum += i
		}
		return sum
	*/
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, 0)
	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	/*
		sums := make([]int, 0)
		for _, arr := range numbers {
			if len(arr) > 0 {
				tail := arr[1:]
				sums = append(sums, Sum(tail))
			} else {
				sums = append(sums, 0)
			}
		}
		return sums
	*/
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(transactions []Transaction, account Account) Account {
	return Reduce(transactions, applyTransaction, account)
}
func applyTransaction(a Account, t Transaction) Account {
	if t.From == a.Name {
		a.Balance -= t.Sum
	}
	if t.To == a.Name {
		a.Balance += t.Sum
	}
	return a
}

func Find[A any](numbers []A, f func(A) bool) (value A, found bool) {
	for _, v := range numbers {
		if f(v) {
			return v, true
		}
	}
	return
}
