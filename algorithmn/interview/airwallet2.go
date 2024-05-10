package interview

//
///*
//A customer could have multiple accounts at Airwallex, and we generate the logs of balances (T, X) for each account, where T is the timestamp, and X is the balance at T.
//
//We want to build a feature that reports the overall total balance report, combining all of the accounts at each timestamp.
//
//For example, a user has 3 accounts, and we have the following logs:
//- Account 1: (1, 30), (3, 50)//, (4, 80)//, (10, 30)
//- Account 2: (2, 40)//,
//- Account 3: (1, 50)//, (6, 30)
// k x  T
//
//The new feature should generate the following overall balance report:
//(1, 80), (2, 120), (3, 140), (4, 190), (6, 170), (10, 120)
//
//Write an algorithm that implements this feature.
//
//*/
//

import "fmt"

type BalanceLog struct {
	Ts      int
	Balance int
}

func NewBalanceLog(ts int, balance int) *BalanceLog {
	return &BalanceLog{Ts: ts, Balance: balance}
}

func GetBalances(accounts [][]*BalanceLog) []*BalanceLog {
	numOfAccounts := len(accounts)

	if numOfAccounts == 0 {
		return nil
	}
	if numOfAccounts == 1 {
		return accounts[0]
	}

	accountIndex := make([]int, numOfAccounts) //current timestamp handled
	results := make([]*BalanceLog, 0)
	largestTs := findLargestTs(accounts)
	smallestTime := accounts[0][accountIndex[0]].Ts
	currentProcessedTs := 0
	//accounts := [][][]int{
	//	{{1, 30}, {3, 50}, {4, 80}, {10, 30}},
	//	{{2, 40}, {4, 60}},
	//	{{1, 50}, {6, 30}},
	//}
	for currentProcessedTs < largestTs {
		smallestTime = accounts[0][accountIndex[0]].Ts
		for i := 1; i < numOfAccounts; i++ {
			if accounts[i][accountIndex[i]].Ts < smallestTime {
				smallestTime = accounts[i][accountIndex[i]].Ts
			}
		}

		sum := 0
		for i := 0; i < numOfAccounts; i++ {
			currentAccountLogIndex := accountIndex[i]
			if currentAccountLogIndex >= len(accounts[i])-1 {
				// reach the end, skip
				sum = sum + accounts[i][currentAccountLogIndex].Balance
				continue
			}
			if accounts[i][currentAccountLogIndex].Ts > smallestTime {
				continue
			} else if accounts[i][currentAccountLogIndex].Ts == smallestTime {
				sum = sum + accounts[i][currentAccountLogIndex].Balance
			} else {
				// check whether needs to move it to next
				if currentAccountLogIndex+1 >= len(accounts[i])-1 {
					// do not move
					sum = sum + accounts[i][currentAccountLogIndex].Balance
				}

				if accounts[i][currentAccountLogIndex+1].Ts == smallestTime {
					accountIndex[i] = accountIndex[i] + 1
					currentAccountLogIndex = accountIndex[i]
					sum = sum + accounts[i][currentAccountLogIndex].Balance

				}
			}

			//sum = sum + accounts[i][currentAccountLogIndex].Balance
		}
		results = append(results, &BalanceLog{
			Ts:      smallestTime,
			Balance: sum,
		})
		currentProcessedTs = smallestTime
	}

	// IMPLEMENT HERE
	return results // Placeholder
}

func findLargestTs(accounts [][]*BalanceLog) int {
	largest := 0

	for i := 0; i < len(accounts); i++ {
		if accounts[i][len(accounts[i])-1].Ts > largest { // last log ts
			largest = accounts[i][len(accounts[i])-1].Ts
		}
	}
	return largest
}

func testAccounts() {
	accounts := [][][]int{
		{{1, 30}, {3, 50}, {4, 80}, {10, 30}},
		{{2, 40}, {4, 60}},
		{{1, 50}, {6, 30}},
	}

	// Convert int arrays to BalanceLog slices
	balanceLogs := GetBalanceLogs(accounts)
	result := GetBalances(balanceLogs)
	PrintResult(result)
	// (1, 80) (2, 120) (3, 140) (4, 190) (6, 170) (10, 120)

	accounts2 := [][][]int{
		{{1, 30}, {3, 50}, {4, 80}, {10, 30}},
	}

	balanceLogs2 := GetBalanceLogs(accounts2)
	result2 := GetBalances(balanceLogs2)
	PrintResult(result2)
	// (1, 30) (3, 50) (4, 80) (10, 30)

	accounts3 := [][][]int{
		{{1, 30}, {3, 50}, {5, 80}, {7, 30}},
		{{2, 30}, {4, 0}, {6, 80}, {7, 0}},
	}

	balanceLogs3 := GetBalanceLogs(accounts3)
	result3 := GetBalances(balanceLogs3)
	PrintResult(result3)
	// (1, 30) (2, 60) (3, 80) (4, 50) (5, 80) (6, 160) (7, 30)
}

func PrintResult(res []*BalanceLog) {
	for _, log := range res {
		fmt.Printf("(%d, %d) ", log.Ts, log.Balance)
	}
	fmt.Print("\n")
}

func GetBalanceLogs(accounts [][][]int) [][]*BalanceLog {
	result := make([][]*BalanceLog, len(accounts))
	for i, account := range accounts {
		aList := make([]*BalanceLog, len(account))
		for j, log := range account {
			aList[j] = NewBalanceLog(log[0], log[1])
		}
		result[i] = aList
	}
	return result
}

//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("not success, not enough money")
//
//}
