package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"impl"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//
//作者：不智鱼
//链接：https://www.jianshu.com/p/24ede9e90490
//来源：简书
//简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。

const base_format = "2006-01-02 15:04:05"

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello"))
	fmt.Printf("http.ResponseWriter:%v \n", w)
	fmt.Printf("*http.Request:%v \n", r)
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	//一个客户端一个连接
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	fmt.Printf("wsConn:%v \n", wsConn)

	// 启动线程，不断发消息
	go func() {
		var (
			err error
		)
		Goid()
		for {
			currentTime := time.Now()
			if err = conn.WriteMessage([]byte("heartbeat" + currentTime.Format(base_format))); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}

func main() {

	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
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
