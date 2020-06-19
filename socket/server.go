package socket

import (
	"log"
	"net"
	"os"
)

func Server() {
	// 监听
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("listen err:", err.Error())
		os.Exit(0)
	}
	defer l.Close()
	// 客户端链接请求
	c, err := l.Accept()
	if err != nil {
		log.Println("Accept err:", err.Error())
		os.Exit(0)
	}
	defer c.Close()
	// 接收客户端请求数据
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		log.Println("Read err:", err.Error())
		os.Exit(0)
	}
	log.Println("接收到客户端数据：", string(buf[:n]))
	// 返回客户端数据
	c.Write([]byte("服务端已收到请求信息"))
}
