package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var connectCount int

func connectToService() interface{} {
	connectCount++
	time.Sleep(1 * time.Second)
	return make([]int, 1)
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		s := p.New()
		log.Printf("Put connection %d into the pool %p\n", i, s)
		p.Put(s)
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
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			log.Printf("get connection from the pool %p\n", svcConn)
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			log.Printf("put connection back into the pool %p\n", svcConn)
			conn.Close()
		}
	}()
	return &wg
}

func init() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}
