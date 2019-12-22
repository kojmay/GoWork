package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c,  c.RemoteAddr(), "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,  c.RemoteAddr(), "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c,  c.RemoteAddr(), "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 2*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening ...")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		fmt.Println("\t connection start:", conn.RemoteAddr())
		go handleConn(conn)
	}
}