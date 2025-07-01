package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var ErrNotConnected = fmt.Errorf("not connected")

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

	return nil
}

func (c *GoTelnetClient) Close() error {
	if c.conn == nil {
		return ErrNotConnected
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *GoTelnetClient) Send() error {
	if c.conn == nil {
		return ErrNotConnected
	}
	return processIO(c.in, c.conn)
}

func (c *GoTelnetClient) Receive() error {
	if c.conn == nil {
		return ErrNotConnected
	}
	return processIO(c.conn, c.out)
}

func processIO(src io.Reader, dst io.Writer) error {
	buf := bufio.NewScanner(src)
	for buf.Scan() {
		_, err := dst.Write([]byte(fmt.Sprintf("%s\n", buf.Text())))
		if err != nil {
			return err
		}
	}
	return nil
}
