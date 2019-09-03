# Account
Simulation of bank account transactions

## Overview
Account is a package that simulates bank account transactions. It supports the following features: 
- Opening of new account
- Depositing funds to an account
- Checking balance of an account
- Withdrawal of funds from an account
- Closing of an account

## API Usage 
```go
    // Opens (creates) a new account
    user := account.Open(200)

    // Deposits the amount given and returns the amount deposited
    _, err := user.Deposit(100)
	if err != nil {
		fmt.Println(err)
    }
    
    // Gets the balance of the account
    bal, err := user.Balance()
	if err != nil {
		fmt.Println(err)
    }
    
    fmt.Println(bal)
    // Returns 300

    // Withdraws funds from the account. It returns amount withdrawn
    _, err = user.Withdraw(50)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
    
    fmt.Println(user.Bal)
    // Returns 250

    // Closes all activities on that account. Returns balance left
    payout, err := user.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
    
    // Withdrawing funds from a closed account returns an error
    _, err = user.Withdraw(100)
	if err != nil {
		fmt.Println(err)
        // "Cannot perform transactions a on closed account"
    }
```


## Todos
- Loan
- Refactor
- Make a web interface
- Integrate with database