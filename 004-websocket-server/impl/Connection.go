package impl

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
	uuid      string
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	//var (
	//	cid uuid.UUID
	//)
	cid, _ := uuid.NewUUID()

	conn = &Connection{
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		wsConn:    wsConn,
		closeChan: make(chan byte, 1),
		uuid:      cid.String(),
	}
	go conn.readLoop()
	go conn.writeLoop()
	return conn, nil
}

//获取连接的编号
func (conn *Connection) GetCID() string {
	return conn.uuid
}

func (conn *Connection) HasClose() bool {
	return conn.isClosed
}

//读取消息
func (conn *Connection) ReadMessage() (data []byte, err error) {

	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("Connection is closed")
	}
	//data= <- conn.inChan
	//return data,nil
	return
}

//发送消息
func (conn *Connection) WriteMessage(data []byte) (err error) {

	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("Connection is closed")
	}
	return

}

func (conn *Connection) Close() {
	//线程安全的Close(可重入)
	conn.wsConn.Close()
	//conn.closeChan<-1
	//加锁同步，保证closeChan只被关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()

}

//循环读取websocket的消息
func (conn *Connection) readLoop() {
	var (
		err  error
		data []byte
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR

		}
		//阻塞在这里，等待inChan有空闲的空间
		conn.inChan <- data
	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {

		select {
		case data = <-conn.outChan:

		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}

	}
ERR:
	conn.Close()

}
