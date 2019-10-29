package main

import (
	"fmt"
	"net"
)

func process2(conn net.Conn) {
	//这时需要延时关闭conn
	defer conn.Close()

	processer := &Processer{
		Conn: conn,
	}
	err := processer.control()
	if err != nil {
		fmt.Println("客户端和服务器端的通信协程错误 err=", err)
		return
	}
}

func main() {
	fmt.Println("服务器[新结构]在8889端口监听....")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("liste Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端端保持通讯
		go process2(conn)
	}
}
