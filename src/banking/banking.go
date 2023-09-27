package banking

import (
	"fmt"
	"errors"
)

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
		er = errors.New("Insufficient funds.")
	} else {
		w.balance -= amount
	}

	return
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}