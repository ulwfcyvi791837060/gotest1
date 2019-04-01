// +build ignore

package main

import (
	"fmt"
	"time"
)

//Administrator@DESKTOP-BD34GOG MINGW64 /f/workspaceApi/gotest1/src/main

//go run -race race.go

func main()  {

	a := 1
	go func(){
		a = 2
	}()
	a = 3
	fmt.Println("a is ", a)

	time.Sleep(2 * time.Second)


}