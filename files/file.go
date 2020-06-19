package files

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// ReadFile
// ioutil.ReadFile(filename string) 调用了ioutil.readAll()
// ioutil.readAll()对os.Open os.Close 封装实现
func ReadFile(filename string) (s string, err error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	s = string(f)
	return
}

// ReadAll
// 和ReadFile一样，同样是调用了ioutil.readAll()
func ReadAll(filename string) (s string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	s = string(b)
	return
}

// Read
func Read(filename string) (s string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	var n int64 = bytes.MinRead
	if fi, err := f.Stat(); err == nil {
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}
	}
	var buf bytes.Buffer
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(n)) == n {
		buf.Grow(int(n))
	}
	_, err = buf.ReadFrom(f)
	s = string(buf.Bytes())
	return
}

// OpenRead
func OpenRead(filename string) (s string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	s = string(buf)
	return
}

// WriteFile
func WriteFile(filename string, s string) error {
	return ioutil.WriteFile(filename, []byte(s), os.ModePerm)
}

// Write
func Write(filename string, s string) (err error) {
	f, err := os.OpenFile(filename ,os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write([]byte(s))
	return
}





