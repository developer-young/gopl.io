// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"sync"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	input := bufio.NewScanner(c)
	timer := time.NewTimer(time.Duration(10) * time.Second)
	renew := make(chan struct{})

	go func() {
		for {
			select {
				case <-timer.C:
					fmt.Println("connect timeout 10s")
					c.Close()
					return
				case <-renew:
					fmt.Println("timer has reset")
					timer.Reset(time.Duration(10) * time.Second)
			}
		}
	} ()

	for input.Scan() {
		wg.Add(1)
		go func() {
			renew <- struct{}{}
			echo(c, input.Text(), 1*time.Second, wg)
		} ()

	}

	wg.Wait()
	if c, ok := c.(*net.TCPConn); ok {
		c.CloseWrite()
	} else {
		c.Close()
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
