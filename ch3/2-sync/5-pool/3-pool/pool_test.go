package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		begin := time.Now()
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		fmt.Printf("time to connect is %s\n", time.Since(begin))
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
