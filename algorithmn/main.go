package main

import (
	"fmt"
	"interview/algorithmn/interview"
)

func main() {
	accountList := make([]interview.Account, 0)
	account1 := interview.Account{
		"a",
		70,
	}

	account2 := interview.Account{
		"b",
		70,
	}

	account3 := interview.Account{
		"a",
		200,
	}

	accountList = append(accountList, account1)
	accountList = append(accountList, account2)
	accountList = append(accountList, account3)

	fmt.Println(interview.MakeTransfer(accountList, 100))

}
