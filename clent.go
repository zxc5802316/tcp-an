package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//Name	            Length	        			Remork
//Package Length	4 bytes	        			header + body length
//Header Length	    2 bytes	        			protocol header length
//Protocol Version	2 bytes	 					protocol version
//Operation			4 bytes						operation for request
//Sequence id		4 bytes						sequence number chosen by client
//Body				PackLen + HeaderLen	 		binary body bytes



func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	fmt.Println("Please input data...")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err:", err)
			return
		}
		//编码数据
		conn.Write(encode(input))
		_, err = conn.Read(buf)
		if err != nil {

			return
		}
		decode(buf)
	}

}
