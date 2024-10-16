package microservices

import (
	"bakend-settings/common"
	"fmt"
	"net"
	// "time"
)

type TcpServer struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan []byte
}

func NewTcpServer(listenAddr string) *TcpServer {
	return &TcpServer{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan []byte, 10),
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

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			break // Exit loop on error
		}
		// Log raw received data
		fmt.Println("Raw data received:", string(buf[:n]))

		// Process the received message
		message, convErr := common.Convert(string(buf[:n]))
		if convErr != nil {
            fmt.Println("Error converting message:", convErr)
            break // Exit the read loop or handle as needed
        }
		  // Check if the message is nil
		  if message == nil {
            fmt.Println("Received nil message")
            break // Handle nil message case
        }

		fmt.Println("Message received:", message)

		// Send a response back to the client (NestJS service)
		// Construct a proper JSON response
		// Construct a proper JSON response
		response, err := common.ConvertToString(message.Pattern, message.ID, message.Data)
		if err != nil {
			fmt.Println("Error constructing response:", err)
		}
		fmt.Println("Response is :", response)

		_, writeErr := conn.Write([]byte(response))
		if writeErr != nil {
			fmt.Println("Write error:", writeErr)
			break
		}
	}
}

// func main() {
// 	config.LoadEnv()
// 	fmt.Printf("server is running on %v\n", config.GetEnv("TCP_SERVER"))
// 	server := NewServer(config.GetEnv("TCP_SERVER"))
// 	log.Fatal(server.Start())
// }
