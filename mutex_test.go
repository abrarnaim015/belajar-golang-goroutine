package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T)  {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}
	
	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (accont *BankAccount) AddBalance(amount int)  {
	accont.RWMutex.Lock()
	accont.Balance = accont.Balance + amount
	accont.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int  {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestReadWriteMutex(t *testing.T)  {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock()  {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock()  {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int)  {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int)  {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T)  {
	user1 := UserBalance {
		Name: "Abrar",
		Balance: 1000000,
	}
	user2 := UserBalance {
		Name: "Naim",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)


	time.Sleep(5 * time.Second)
}