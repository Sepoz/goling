package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin new type
type Bitcoin int

// Stringer interface
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit Wallet method
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds variable
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw method
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance Wallet method
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
