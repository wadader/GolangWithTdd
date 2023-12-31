package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

var insuffucientBalanceErr = errors.New("insufficient funds, withdrawn greater than balance")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return insuffucientBalanceErr
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin { return w.balance }
