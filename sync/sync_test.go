package sync

import (
	"reflect"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertEqual(t, 3, counter)
	})
	t.Run("safely run concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertEqual(t, wantedCount, counter)
	})
}

func assertEqual(t *testing.T, expected int, got *Counter) {
	t.Helper()
	if !reflect.DeepEqual(expected, got.Value()) {
		t.Errorf("got %v, want %v", got.Value(), expected)
	}
}
