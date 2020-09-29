package main

import (
	"fmt"

	"github.com/XxCtrlZxX/learngo/lecture2/accounts"
)

func main() {
	account := accounts.NewAccount("CtrlZ")
	fmt.Println(account)
}
