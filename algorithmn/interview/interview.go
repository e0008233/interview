package interview

//Say we have a number of bank accounts, and we want To make some transfers between these
//accounts at the end of each day To make sure that each Account will have at least 100 dollars after these transfers.
//Please write a program To implement this. Each Transfer will be presented as a tuple of (from_account, to_account, Amount). For example, we have:
//
//Example :
//bank_accounts = [(“a”, 80), (“b”, 180)]
//transfers = make_transfers(bank_accounts)
//# transfers = [(“b”, “a”, 20)]

// n accounts, money >= 100 * n

type Transfer struct {
	From   string
	To     string
	Amount int
}

type Account struct {
	Name    string
	Balance int
}

// 	140, 70, 160, 60
func MakeTransfer(accountList []Account, minBalance int) ([]Transfer, bool) {

	// check on the accounts

	maxIndex := -1 // first element > minBalance
	result := make([]Transfer, 0)
	// scan through Account list

	for index, _ := range accountList {
		if accountList[index].Balance < minBalance {
			// handle the min Account
			differenceNeeded := minBalance - accountList[index].Balance
			for differenceNeeded > 0 {
				for maxIndex <= len(accountList)-1 && (maxIndex == -1 || accountList[maxIndex].Balance <= minBalance) {
					maxIndex++
				}
				if maxIndex >= len(accountList) {
					// cannot find the next max To make the Transfer
					return nil, false
				}
				maxTransferPossible := accountList[maxIndex].Balance - minBalance
				if maxTransferPossible > differenceNeeded {
					// one Transfer is enough
					newTransfer := Transfer{
						From:   accountList[maxIndex].Name,
						To:     accountList[index].Name,
						Amount: differenceNeeded,
					}
					// update the Account
					accountList[index].Balance = accountList[index].Balance + differenceNeeded
					accountList[maxIndex].Balance = accountList[maxIndex].Balance - differenceNeeded
					differenceNeeded = 0
					result = append(result, newTransfer)
				} else {
					// one Transfer is not enough
					newTransfer := Transfer{
						From:   accountList[maxIndex].Name,
						To:     accountList[index].Name,
						Amount: maxTransferPossible,
					}
					// update the Account
					accountList[index].Balance = accountList[index].Balance + maxTransferPossible
					accountList[maxIndex].Balance = accountList[maxIndex].Balance - maxTransferPossible
					differenceNeeded = differenceNeeded - maxTransferPossible
					result = append(result, newTransfer)
				}

			}
		}
	}
	return result, true
}
