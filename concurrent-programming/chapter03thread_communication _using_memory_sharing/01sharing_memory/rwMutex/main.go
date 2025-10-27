package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	data = make(map[string]string)
	rwMu sync.RWMutex
	wg   sync.WaitGroup
)

func read(key string) string {
	rwMu.RLock()
	defer rwMu.RUnlock()
	return data[key]
}

func write(key, value string) {
	rwMu.Lock()
	data[key] = value
	rwMu.Unlock()
}

func main() {
	// สร้าง goroutine เขียนข้อมูล
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("worker-%d", id)
			value := fmt.Sprintf("value-%d", id)
			write(key, value)
			fmt.Println("Write : ", key, " = ", value)
		}(i)
	}

	// รอให้เขียนเสร็จก่อน แล้วค่อยอ่าน
	time.Sleep(500 * time.Millisecond)

	// สร้าง goroutine อ่านข้อมูล
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("worker-%d", id)
			val := read(key)
			fmt.Println("Read : ", key, " = ", val)
		}(i)
	}
	wg.Wait()
	fmt.Println("Final Data : ", data)

}
