package pointers

import "errors"

type Wallet struct {
	balance float64
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid amount to deposit")
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return errors.New("Insufficient balance")
	}

	w.balance -= amount
	return nil
}
