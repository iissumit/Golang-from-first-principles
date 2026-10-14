package main

import "fmt"

func main() {
	account := BankAccount{
		Owner:   "Alex",
		Balance: 5000,
	}
	fmt.Println(account.Owner)
	fmt.Println(account.Balance)

	account.Balance = 6000
	fmt.Println(account.Balance)
}

// BankAccount is struct nts a name of our new type
// struct tells Go were grouping fields together
// Owner is a fielf of type string
// Balance is a field of type integer
type BankAccount struct {
	Owner   string
	Balance int
}
