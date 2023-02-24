package safemap

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeMap(t *testing.T) {

	store := New()

	var waiter sync.WaitGroup

	waiter.Add(1)

	go func() {
		for i := 0; i < 100; i++ {
			store.Insert(fmt.Sprintf("0x%04X", i), i)
			if i > 0 && i%15 == 0 {
				fmt.Printf("Inserted %d items\n", store.Len())
			}
		}
		fmt.Printf("Inserted %d items\n", store.Len())
		waiter.Done()
	}()

	waiter.Wait()
	fmt.Printf("Finished with %d items\n", store.Len())

	data := store.Close()
	fmt.Println("Closed")
	fmt.Printf("len == %d\n", len(data))
}
