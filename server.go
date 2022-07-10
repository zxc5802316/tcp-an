package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	fmt.Println("Listening....")
	fmt.Println()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		decode(line)
		fmt.Println("-----------------------------------------------")
		wr.Write(encode("Receive Success!"))
		wr.Flush()
	}
}