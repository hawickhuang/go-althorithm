package syncmap

import "sync"

func readWriteMap() {
	m := make(map[int]int)

	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				mu.RLock()
				_ = m[1]
				mu.RUnlock()
			}
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				mu.Lock()
				m[2] = 2
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
