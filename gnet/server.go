package gnet

import (
	"fmt"
	"net"
)

type Server struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Name:      "服务端",
		IpVersion: "tcp",
		Ip:        ip,
		Port:      port,
	}
}

func (s *Server) Start() {
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("resolve ip_version error: ", err)
		return
	}

	listen, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("listen tcp error: ", err)
		return
	}
	fmt.Printf("start server, server host: %s:%d \n", s.Ip, s.Port)
	for {
		conn, err := listen.Accept()
		fmt.Printf("%p \n", &conn)
		if err != nil {
			fmt.Println("accept conn error: ", err)
			continue
		}

		go func() {
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("read client data", err)
					return
				}
				fmt.Println("accept client message: ", string(buf[:n]))

				_, err = conn.Write(buf[:n])
				if err != nil {
					fmt.Println("write data", err)
				}
			}
		}()
	}
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}
