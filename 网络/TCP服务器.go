package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	lisenter, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer lisenter.Close()
	for {
		conn, err1 := lisenter.Accept()
		if err != nil {
			fmt.Println("err1 = ", err1)
			return
		}
		buf := make([]byte, 1024)
		num, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("err2 = ", err2)
			return
		}
		//fmt.Println("buf =", string(buf[:num]))
		//defer conn.Close()

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect success")

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		fmt.Println("[%s]: %s\n", addr, string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		conn.Close()

	}
}
