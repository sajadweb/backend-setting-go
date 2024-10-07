package main

import (
	"bakend-settings/config"
	"log"
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch    chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	// Start accept loop in a goroutine
	go s.acceptLoop()

	// Block until the server is stopped (via closing quitch)
	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accepted error", err)
			continue
		}
		go s.readLoop(conn)
	}
}
func (s *Server) readLoop(conn net.Conn)  {
	defer conn.Close()

	buf := make([]byte, 2048)

	for{
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}
		
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}
func main() {
	config.LoadEnv()
	fmt.Printf("server is  %v",config.GetEnv("TCP_SERVER"))
	server :=NewServer(config.GetEnv("TCP_SERVER"))
	log.Fatal(server.Start())
}
