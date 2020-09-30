package main

import (
	"fmt"

	"github.com/XxCtrlZxX/lecture2/accounts"
)

func main() {
	account := accounts.NewAccount("CtrlZ")
	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
