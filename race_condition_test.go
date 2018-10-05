package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	b, wg := &bank{balance: 0, m: new(sync.RWMutex)}, new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		b.Deposit(200)
		fmt.Printf("Deposit: %d\n", b.Balance())
		wg.Done()
	}()
	go func() {
		b.Deposit(100)
		wg.Done()
	}()
	wg.Wait()
}
