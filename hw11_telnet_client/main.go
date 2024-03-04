package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var duration time.Duration

func init() {
	flag.DurationVar(&duration, "timeout", 10*time.Second, "timeout")
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
	in := io.NopCloser(bufio.NewReader(os.Stdin))
	out := os.Stdout

	client := MustConnect(host, port, duration, in, out)

	os.Stderr.WriteString(fmt.Sprintf("...Connected to %s:%s\n", host, port))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err := client.Send()
		if err != nil {
			log.Println("connection closed by peer")
			cancel()
		}
	}()

	go func() {
		err := client.Receive()
		if err != nil {
			log.Println()
			cancel()
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
		cancel()
	case <-ctx.Done():
	}
}

func MustConnect(host, port string, duration time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	client := NewTelnetClient(net.JoinHostPort(host, port), duration, in, out)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
