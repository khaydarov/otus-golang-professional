package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &GoTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type GoTelnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (c *GoTelnetClient) Connect() error {
	var err error
	dialer := &net.Dialer{}
	dialer.Timeout = c.timeout
	c.conn, err = dialer.Dial("tcp", c.address)
	if err != nil {
		return err
	}

	clientMessagesChannel := ioReaderChannel(c.in)

	defer c.conn.Close()
	return nil
}

func (c *GoTelnetClient) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *GoTelnetClient) Send() error {
	var err error
	for message := range clientMessagesChannel {
		_, err = c.conn.Write([]byte(fmt.Sprintf("%s\n", message)))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *GoTelnetClient) Receive() error {
	var err error
	buf := bufio.NewScanner(c.conn)
	for buf.Scan() {
		_, err = c.out.Write([]byte(fmt.Sprintf("%s\n", buf.Text())))
		if err != nil {
			return err
		}
	}
	return nil
}

func ioReaderChannel(in io.ReadCloser) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		buf := bufio.NewScanner(in)
		for buf.Scan() {
			out <- buf.Text()
		}
	}()
	return out
}
