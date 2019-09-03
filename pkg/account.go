package account

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

// Account is a blueprint of a user bank account
type Account struct {
	Bal    float64
	Closed bool
	sync.Mutex
}

// Open creates a new acoount for the given user and return an instance of the account type
func Open(bal float64) *Account {
	// Check if Bal is greater than 0
	if bal < 0 {
		return nil
	}
	account := &Account{Bal: bal, Closed: false}
	//return &Account{Bal: bal, closed: false}
	return account
}

// Deposit credits money to an account with the given amount and returns amount deposited and error
func (a *Account) Deposit(amount float64) (float64, error) {
	a.Lock()
	defer a.Unlock()

	var err error

	// No operations for a closed account
	if a.Closed == true {
		err = errors.New("Cannot perform transactions on a closed account")
		return 0, err
	}

	// Deposits must be more than zero
	if amount <= 0 {
		err = errors.New("Deposits must be more than 0")
		return 0, err
	}

	a.Bal = a.Bal + amount

	return amount, nil
}

// Balance returns the Bal of an account
func (a *Account) Balance() (float64, error) {
	a.Lock()
	defer a.Unlock()

	var err error

	// No transactions on a closed account
	if a.Closed == true {
		err = errors.New("Cannot perform transactions a on closed account")
		return 0, err
	}

	return a.Bal, nil
}

// Withdraw takes an amount to deduct from the account and returns the amount deducted
func (a *Account) Withdraw(amount float64) (float64, error) {

	a.Lock()
	defer a.Unlock()

	var err error

	// No transactions on a closed account
	if a.Closed == true {
		err = errors.New("Cannot perform transactions a on closed account")
		return 0, err
	}

	// Withdrawal logic
	if amount > a.Bal {
		err = errors.New("Insufficient funds")
		return 0, err
	}

	if amount < 0 {
		err = errors.New("Cannot withdraw below 0")
		return 0, err
	}

	a.Bal -= (amount)

	return amount, nil
}

// Close closes or terminates the users account and returns all the balance in that account as payout
func (a *Account) Close() (float64, error) {
	a.Lock()
	defer a.Unlock()

	var err error
	var payout float64

	// No transactions on a closed account
	if a.Closed == true {
		err = errors.New("Cannot perform transactions a on closed account")
		return 0, err
	}

	payout = a.Bal
	a.Closed = true
	a.Bal = a.Bal - a.Bal

	return payout, nil
}

// Simulate API Usage
func Simulate() {
	user := Open(2000)
	fmt.Println(user)

	_, err := user.Deposit(100)
	if err != nil {
		fmt.Println(err)
	}

	bal, err := user.Balance()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bal)

	_, err = user.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(user.Bal)

	payout, err := user.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(payout)
	fmt.Println(user)

	_, err = user.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(user.Bal)
}
