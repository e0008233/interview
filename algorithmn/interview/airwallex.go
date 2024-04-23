package interview

import (
	"errors"
	"fmt"
)

/*
Say we have a number of bank accounts, and we want to make some transfers between these accounts
at the end of each day to make sure that each account will have at least 100 dollars after these
transfers.
Please write a program to implement this. Each transfer will be presented as a tuple of (from_account,
to_account, amount).
For example:
bankAccounts = [(“a”, 80), (“b”, 180), ("c", 90)]
transfers = makeTransfers(bankAccounts)
# transfers = [(“b”, “a”, 20), ("b", "c", 10)]
*/

type transfer struct {
	FromAccountName string
	ToAccountName   string
	Amount          int
}

type account struct {
	Name    string
	balance int
}

var targetBalance = 100

func TransferMoney() {
	accounts := make([]account, 0)
	accounts = append(accounts, account{
		Name:    "a",
		balance: 80,
	})
	accounts = append(accounts, account{
		Name:    "b",
		balance: 180,
	})
	accounts = append(accounts, account{
		Name:    "c",
		balance: 60,
	})

	totalMoney := 0
	for _, curr := range accounts {
		totalMoney = totalMoney + curr.balance
	}
	result, err := MakeTransfer(accounts)

	for _, curr := range accounts {
		totalMoney = totalMoney - curr.balance
	}

	if totalMoney != 0 {
		fmt.Printf("transfer check, result: %v\n", totalMoney)
		fmt.Println("fund check failed")
		// todo: alarm
	}
	if err != nil {
		fmt.Println("not success, not enough money")
	} else {
		fmt.Printf("success, result: %v\n", result)
	}

	testcases := MakeTestCases()
	for _, testcase := range testcases {
		fmt.Printf("-------------------------\n")

		fmt.Printf("running test case: %v\n", testcase)
		// check total totalMoney
		// not make transfer if not enough
		result2, err2 := MakeTransfer(testcase)
		if err2 != nil {
			fmt.Println("not success, not enough money")
		} else {
			fmt.Printf("success, result: %v\n", result2)
		}
	}

}

func MakeTransfer(accounts []account) ([]transfer, error) {
	if len(accounts) == 0 {
		return nil, errors.New("empty accounts")
	}
	moreAccountIndex := 0

	transferRecords := make([]transfer, 0)

	for lessAccountIndex := 0; lessAccountIndex < len(accounts); lessAccountIndex++ {
		if accounts[lessAccountIndex].balance < 0 {
			return nil, errors.New("negative value found")
		}
		if accounts[lessAccountIndex].balance < targetBalance {
			// needs to find the money to be transfer in
			diffNeeded := targetBalance - accounts[lessAccountIndex].balance

			for diffNeeded > 0 {
				moreAccountIndex = findNextMoreAccount(accounts, moreAccountIndex)

				if moreAccountIndex == -1 {
					return nil, errors.New("cannot find next more account")
				}
				if accounts[moreAccountIndex].balance-targetBalance >= diffNeeded {
					// enough to make transfer
					newTransfer := transfer{
						FromAccountName: accounts[moreAccountIndex].Name,
						ToAccountName:   accounts[lessAccountIndex].Name,
						Amount:          diffNeeded,
					}
					transferRecords = append(transferRecords, newTransfer)
					// update balance
					accounts[lessAccountIndex].balance = accounts[lessAccountIndex].balance + diffNeeded
					accounts[moreAccountIndex].balance = accounts[moreAccountIndex].balance - diffNeeded
					diffNeeded = 0
				} else {
					diffCanSupport := accounts[moreAccountIndex].balance - targetBalance
					newTransfer := transfer{
						FromAccountName: accounts[moreAccountIndex].Name,
						ToAccountName:   accounts[lessAccountIndex].Name,
						Amount:          diffCanSupport,
					}
					transferRecords = append(transferRecords, newTransfer)
					// update balance
					accounts[lessAccountIndex].balance = accounts[lessAccountIndex].balance + diffCanSupport
					accounts[moreAccountIndex].balance = accounts[lessAccountIndex].balance + diffCanSupport
					diffNeeded = diffNeeded - diffCanSupport
				}
			}

		}

	}

	return transferRecords, nil

}

func findNextMoreAccount(accounts []account, index int) int {
	for i := index; i < len(accounts); i++ {
		if accounts[i].balance > targetBalance {
			return i
		}
	}

	return -1
}

func MakeTestCases() [][]account {
	testAccountList := make([][]account, 0)

	test1 := make([]account, 0)
	test1 = append(test1, createAccount("a", 80))
	test1 = append(test1, createAccount("b", 180))
	test1 = append(test1, createAccount("c", 80))

	test2 := make([]account, 0)
	test2 = append(test2, createAccount("a", 0))
	test2 = append(test2, createAccount("b", 0))
	test2 = append(test2, createAccount("c", 0))

	test3 := make([]account, 0)
	test3 = append(test3, createAccount("a", 120))
	test3 = append(test3, createAccount("b", 120))
	test3 = append(test3, createAccount("c", 120))

	test4 := make([]account, 0)
	test4 = append(test4, createAccount("a", 80))
	test4 = append(test4, createAccount("b", 111))
	test4 = append(test4, createAccount("c", 80))
	testAccountList = append(testAccountList, test1)
	testAccountList = append(testAccountList, test2)
	testAccountList = append(testAccountList, test3)
	testAccountList = append(testAccountList, test4)

	return testAccountList
}

func createAccount(name string, balance int) account {
	return account{
		Name:    name,
		balance: balance,
	}
}
