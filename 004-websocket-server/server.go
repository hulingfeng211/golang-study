package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hulingfeng211/golang-study/004-websocket-server/impl"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	onlineConnection = make(map[string]*impl.Connection)
	mutex            sync.Mutex
)

//登记已经打开的连接
func registryConnection(con *impl.Connection) {
	onlineConnection[con.GetCID()] = con

}

//删除已经关闭的连接
func removeConnection(cid string) {
	mutex.Lock()
	//if conn,ok:=onlineConnection[cid];ok {
	//	conn.Close()
	//}
	delete(onlineConnection, cid)
	mutex.Unlock()
}

func main() {
	onlineConnection = make(map[string]*impl.Connection)
	//websocket.Upgrade
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		var (
			wsConn *websocket.Conn
			err    error
			//msgType int
			data []byte
			conn *impl.Connection
		)
		if wsConn, err = upgrader.Upgrade(writer, request, nil); err != nil {
			return
		}
		if conn, err = impl.InitConnection(wsConn); err != nil {
			goto ERR
		}
		registryConnection(conn)
		//在线连接的维护
		go func() {
			for {
				fmt.Println(len(onlineConnection))
				for cid, conn := range onlineConnection {
					if conn.HasClose() {
						removeConnection(cid)
					}
				}
				time.Sleep(1 * time.Second)
			}

		}()

		//心跳检测的协程
		go func() {
			var (
				err error
			)
			for {
				log.Print(len(onlineConnection))
				if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
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
		//ERROR:
		//	conn.Close()
	ERR:
		conn.Close()
		//todo 关闭连接

	})
	http.ListenAndServe(":7777", nil) //会阻塞

	fmt.Println("Hello,Web Socket")

}
