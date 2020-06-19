package logs

import (
	"log"
	"os"
)

// Log
func Log(filename string, s string) (err error) {
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	// 将文件设置为log输出的文件
	log.SetOutput(f)
	// 输出前缀
	log.SetPrefix("[log]")
	// log格式
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	// log输出到文件
	log.Println([]string{s})
	return
}

// Logger
func Logger(filename string, s string) (err error) {
	f, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	logger := log.New(f, "[logger]", log.LstdFlags | log.Lshortfile | log.LUTC)
	// log输出到文件
	logger.Println([]string{s})
	return
}
