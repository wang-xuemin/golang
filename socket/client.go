package socket

import (
	"log"
	"net"
	"os"
)

func Client() {
	// 发起请求链接
	c, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("Dial err:", err.Error())
		os.Exit(0)
	}
	defer c.Close()
	// 发送数据
	_, err = c.Write([]byte("are you ready"))
	if err != nil {
		log.Println("Write err:", err.Error())
		os.Exit(0)
	}
	// 读取返回数据
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		log.Println("Read err:", err.Error())
		os.Exit(0)
	}
	log.Println("接收到服务器返回数据：", string(buf[:n]))
}
