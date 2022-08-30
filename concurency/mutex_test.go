package concurency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) ReadBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				account.AddBalance(1)
				fmt.Println(account.ReadBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance : ", account.Balance)
}
