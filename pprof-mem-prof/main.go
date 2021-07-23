package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func go_ticker() {
	fmt.Println("Started go_ticker")
	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
	fmt.Println("Finished go_ticker")
}

func norm_ticker() {
	fmt.Println("Started norm_ticker")
	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
	fmt.Println("Finished norm_ticker")
}

func norm_10() {
	var buf []byte
	fmt.Println("Started norm_10")
	for i := 0; i < 10; i++ {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
	fmt.Println("Finished norm_10")
}

func iter() {
	fmt.Println("Started iter")
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration : ", i)
	}
	norm_ticker()
	fmt.Println("Finished iter")
}

// Run for a period of time: fatal error: runtime: out of memory
func main() {
	fmt.Println("Started main")
	// Open pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()
	go go_ticker()
	norm_10()
	iter()
	fmt.Println("Finished main")
}
