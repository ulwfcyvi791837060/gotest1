package test

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func Test0001(test *testing.T) {
	closeChan := make(chan byte, 1)
	close(closeChan)
	select {
	case <-closeChan:
		fmt.Println("TestDecrypt = ")
	}
}

func Test0002(test *testing.T) {
	for i := 0; i < 20; i++ {
		go func() {
			Goid()
			for {
				b := make([]byte, 10)
				os.Stdin.Read(b) // will block
			}
		}()
	}
	select {}
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic recover:panic info:%v \n", err)
		}
	}()
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Printf("get goroutine id: %v \n", id)
	return id
}
