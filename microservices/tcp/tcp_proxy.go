package tcp

import (
	"bakend-settings/microservices/tcp/common"
	"fmt"
	"net"
	// "time"
)

type TcpServer struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan []byte
	routes     map[string]func(req *common.TcpRequest)
}

func NewTcpServer(listenAddr string) *TcpServer {
	return &TcpServer{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan []byte, 10),
		routes:     make(map[string]func(req *common.TcpRequest)),
	}
}

func (s *TcpServer) Start() error {
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
	close(s.msgch)
	return nil
}

func (s *TcpServer) acceptLoop() {
	fmt.Println("Accepted start")
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accepted error", err)
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *TcpServer) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	fmt.Println("Read readLoop:")
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			break // Exit loop on error
		}
		// Log raw received data
		//fmt.Println("Raw data received:", string(buf[:n]))

		// Process the received message
		request, convErr := common.MakeTcpRequest(conn, string(buf[:n]))
		if convErr != nil {
			fmt.Println("Error converting message:", convErr)
			break // Exit the read loop or handle as needed
		}
		// Check if the message is nil
		if request == nil {
			fmt.Println("Received nil message")
			break // Handle nil message case
		}
		//fmt.Printf("pattern is pattern %v", request.Pattern)
		// Check if the pattern exists in the routes map and call the handler if it does
		handler, exists := s.routes[request.Pattern];
		if exists {
			go handler(request)
		} else {
			fmt.Printf("No handler found for pattern: %v\n", request.Pattern)
		}
		 
	}
}
func (s *TcpServer) Pattern(pattern string, _func func(req *common.TcpRequest)) {
	s.routes[pattern] = _func
}

func Serve(ip string) *TcpServer { 
	return  NewTcpServer(ip) 
}
