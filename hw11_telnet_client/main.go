package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var duration time.Duration

func init() {
	flag.DurationVar(&duration, "timeout", 0, "timeout")
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("host and port are required")
		return
	}

	host := args[0]
	port := args[1]

	address := net.JoinHostPort(host, port)

	in := bufio.NewReader(os.Stdin)

	client := NewTelnetClient(address, duration, in, os.Stdout)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
	}

	os.Stderr.WriteString(fmt.Sprintf("...Connected to %s\n", address))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		client.Send()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		client.Receive()
	}()

	wg.Wait()
}
