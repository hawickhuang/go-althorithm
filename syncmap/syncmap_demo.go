package syncmap

import (
	"fmt"
	"sync"
)

func syncMapReadWrite() {
	var sm sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Store("key1", "value1")
			sm.Store(1, 1)
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if v, ok := sm.Load("key1"); ok {
				fmt.Printf("load key1=%s\n", v)
			}
			if v, ok := sm.Load(1); ok {
				fmt.Printf("load 1=%d\n", v)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if v, ok := sm.LoadOrStore("key1", "value_1"); ok {
				fmt.Printf("LoadOrStore key1=%s\n", v)
			}
			if v, ok := sm.LoadOrStore("2", "2_2"); ok {
				fmt.Printf("LoadOrStore 2=%d\n", v)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Range(func(key, value interface{}) bool {
				fmt.Printf("key[%v]=value[%v]\n", key, value)
				return true
			})
		}()
	}
	wg.Wait()
}
