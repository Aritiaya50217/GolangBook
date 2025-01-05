package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func pool() {
	// สร้าง sync.Pool และกำหนดฟังก์ชัน New สำหรับสร้างออบเจ็กต์ใหม่
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return "new object"
		},
	}
	// ดึง object จาก pool
	obj1 := myPool.Get().(string)
	fmt.Println("Obj1 : ", obj1)

	// คืน object กลับไปยัง pool
	myPool.Put("reused object")

	// ดึง object จาก myPool อีกครั้ง
	obj2 := myPool.Get().(string)
	fmt.Println("Obj2 : ", obj2)

	// ดึง ​object อีกครั้ง (กรณีไม่มีใน myPool จะสร้างใหม่)
	obj3 := myPool.Get().(string)
	fmt.Println("Obj3 : ", obj3)
}

func calsPool() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// send the pool with 4KB (1KB = 1024 ตัวอักษร)
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}

func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDomain() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost : 8080")
		if err != nil {
			log.Fatal("cannot listen : ", err)
		}
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection : %v", err)
				continue
			}
			connectToService()
			fmt.Println(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

// func init() {
// 	daemonStarted := startNetworkDomain()
// 	daemonStarted.Wait()
// }

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost: 8080")
		if err != nil {
			b.Fatalf("cannot dial host : %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read : %v", err)
		}
		conn.Close()
	}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection : %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func main() {
	fmt.Println("----- pool -----")
	pool()
	fmt.Println("----- calsPool -----")
	calsPool()
	fmt.Println("----- start network daemon -----")
	startNetworkDaemon()
}
