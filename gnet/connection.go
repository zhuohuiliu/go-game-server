package gnet

import (
	"fmt"
	"go-game-server/gface"
	"net"
	"strings"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    int
	isClose   bool
	handleAPI gface.HandleFunc
	ExitChan  chan bool
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine Start...")
	defer func() {
		fmt.Println("Reader is Exit Remote Addr is:", c.Conn.RemoteAddr().String())
		c.Stop()
	}()

	for {
		buf := make([]byte, 1024)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("buffer read error", err)
			continue
		}
		fmt.Println(strings.Contains(string(buf), "quit"))
		if strings.Contains(string(buf), "quit") {
			fmt.Printf("ConnID-%d close\n", c.ConnID)
			break
		}

		err = c.handleAPI(c.Conn, buf, n)
		if err != nil {
			fmt.Println("handleAPI error ConnID:", c.ConnID)
			break
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ConnID:", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop() ConnID:", c.ConnID)
	if c.isClose {
		return
	}
	c.isClose = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnectionID() int {
	return c.ConnID
}

func (c *Connection) RemoteAddr() string {
	return c.Conn.RemoteAddr().String()
}

func (c *Connection) Send(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connID int, callBack_api gface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClose:   false,
		handleAPI: callBack_api,
		ExitChan:  make(chan bool, 1),
	}
	return c
}
