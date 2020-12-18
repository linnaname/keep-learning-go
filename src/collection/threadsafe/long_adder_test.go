package threadsafe

import (
	"sync"
	"testing"
)

func TestLongAdder_Sum(t *testing.T) {
	adder := New()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				adder.Increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	tmp := int64(10) * int64(10)
	if adder.Sum() != tmp || adder.SumThenReset() != tmp || adder.Sum() != 0 {
		t.Errorf("LongAdder logic is wrong")
	}
}
