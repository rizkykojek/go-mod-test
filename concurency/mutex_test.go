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

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
	fmt.Println("Lock ", user.Name)

}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
	fmt.Println("UnLock ", user.Name)
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	user1.Change(-amount)

	user2.Lock()
	user2.Change(amount)

	user2.Unlock()
	user1.Unlock()

	fmt.Println("Transfer Done...")
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rizky",
		Balance: 100,
	}

	user2 := UserBalance{
		Name:    "Kojek",
		Balance: 100,
	}

	go Transfer(&user1, &user2, 8)
	go Transfer(&user2, &user1, 12)

	time.Sleep(6 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance : ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance : ", user2.Balance)
}
