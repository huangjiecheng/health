package main

import (
	"fmt"
	"health/detector"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.RWMutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount
	fmt.Printf("Deposited %d. New balance: %d\n", amount, b.balance)
}

func (b *BankAccount) GetBalance() int {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.balance
}

func main() {
	detector.Start()
	//account := BankAccount{balance: 1000}
	//
	//// Simulate multiple users reading and updating the account balance concurrently
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Printf("Current balance: %d\n", account.GetBalance())
	//	}()
	//}
	//
	//for i := 0; i < 2; i++ {
	//	go func() {
	//		account.Deposit(100)
	//	}()
	//}
	//
	//time.Sleep(time.Second) // Wait for all goroutines to finish
	//
	//fmt.Printf("Final balance: %d\n", account.GetBalance())
}
