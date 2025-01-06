package pointersErrors

import (
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type WalletInterface interface {
	Deposit(amount Bitcoin)
	Balance() Bitcoin
	Withdraw(amount Bitcoin)
}
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

const (
	ErrInsufficientFunds = "cannot withdraw from zero balance to negative amount"
)

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return fmt.Errorf(ErrInsufficientFunds)
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
