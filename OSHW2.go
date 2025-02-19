package main

// Needed libraries
import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Makes sure that the  threads acquire the lock in the order they requested it
type TicketLock struct {
	ticket uint32
	turn   uint32
}

func (t *TicketLock) Lock() {
	myTurn := atomic.AddUint32(&t.ticket, 1) - 1 // Gets a unique ticket for each thread
	for atomic.LoadUint32(&t.turn) != myTurn {
		// Wait until it's this thread's turn
	}
}

func (t *TicketLock) Unlock() {
	atomic.AddUint32(&t.turn, 1)
}

// The compare and swap fucntion
type CASLock struct {
	flag uint32
}

// Keeps spinning until successful
func (c *CASLock) Lock() {
	for !atomic.CompareAndSwapUint32(&c.flag, 0, 1) {

	}
}

func (c *CASLock) Unlock() {
	atomic.StoreUint32(&c.flag, 0) // Releases the lock
}

// Benchmark function: runs multiple goroutines using the given lock
func benchmark(lock sync.Locker, threads int) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			time.Sleep(1 * time.Millisecond)
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func main() {
	// Test with different numbers of threads
	for _, threads := range []int{2, 4, 8, 16} {
		ticketTime := benchmark(&TicketLock{}, threads)
		casTime := benchmark(&CASLock{}, threads)

		fmt.Printf("Threads: %d | Ticket Lock: %v | CAS Lock: %v\n", threads, ticketTime, casTime)
	}
}
