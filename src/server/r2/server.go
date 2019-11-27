package main

import (
	"./impl"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)
	//msgType=0
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heart beat")); err != nil {
				return
			}
			time.Sleep(5 * time.Second)
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
	//TODO:close wsConn
	conn.Close()
}
func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":7777", nil)
}

/*func main() {
	var data string
	s := make(chan string, 10)
	t := make(chan string,1)
	message:=[]string{"1","2","3","4","5"}

	for _,d:=range message{
		go func() {
			s <- "s" +d
		}()

	}
	go func() {
		t<-"data1"
		//close(t)
	}()
	for i:=0;i<1;i++ {
		select {
		case data = <-s:
			fmt.Println(data)
		case <-t:
			fmt.Println("t closed")
		}
	}
}*/
