package main

import (
	"fmt"
	"health/demo"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
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
func generateFlameGraph() {
	// 触发火焰图生成
	resp, err := http.Get("http://localhost:6060/debug/pprof/profile")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	demo.Three()

	//ticker.Ticker(1, func() {
	//	fmt.Println(fmt.Sprintf("NumGoroutine【%d】NumCPU【%d】NumCgoCall【%d】MemProfileRate【%d】",
	//		runtime.NumGoroutine(), runtime.NumCPU(), runtime.NumCgoCall()), runtime.MemProfileRate)
	//})
	//detector.Start()
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

	//go startProfileServer("127.0.0.1:18405")

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	select {}
}

func startProfileServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	fmt.Println(fmt.Sprintf("go_pprof listener on %s", addr))
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(fmt.Sprintf("start profile server failed:%s", err.Error()))
	}
}
