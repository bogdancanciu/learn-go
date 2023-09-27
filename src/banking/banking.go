package banking

import (
	"fmt"
	"errors"
)

var ErrInsufficientFunds = errors.New("Insufficient funds.")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) (er error) {
	if amount > w.balance {
		er = ErrInsufficientFunds
	} else {
		w.balance -= amount
	}

	return
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}