package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-category-app/models"
	"go-category-app/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net"
	"strings"
)

type TCPService struct {
	Repo *repositories.CategoryRepository
}

func NewTCPService(repo *repositories.CategoryRepository) *TCPService {
	return &TCPService{Repo: repo}
}

func (s *TCPService) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error starting TCP server:", err)
	}
	defer listener.Close()

	fmt.Println("TCP server started on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *TCPService) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Connection closed by client:", err)
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		response := s.handleRequest(message)
		conn.Write([]byte(response + "\n"))
	}
}

func (s *TCPService) handleRequest(message string) string {
	parts := strings.Split(message, " ")
	command := parts[0]

	switch command {
	case "GET_ALL":
		categories, err := s.Repo.GetAll()
		if err != nil {
			return "Error: " + err.Error()
		}
		response, _ := json.Marshal(categories)
		return string(response)

	case "GET":
		if len(parts) < 2 {
			return "Error: Missing ID"
		}
		id := parts[1]
		category, err := s.Repo.GetByID(id)
		if err != nil {
			return "Error: " + err.Error()
		}
		response, _ := json.Marshal(category)
		return string(response)

	case "CREATE":
		if len(parts) < 3 {
			return "Error: Missing parameters"
		}
		name := parts[1]
		icon := parts[2]
		category := models.Category{Name: name, Icon: icon}
		_, err := s.Repo.Create(category)
		if err != nil {
			return "Error: " + err.Error()
		}
		return "Category created successfully"

	case "UPDATE":
		if len(parts) < 4 {
			return "Error: Missing parameters"
		}
		id := parts[1]
		name := parts[2]
		icon := parts[3]
		updateData := bson.M{"name": name, "icon": icon}
		_, err := s.Repo.Update(id, updateData)
		if err != nil {
			return "Error: " + err.Error()
		}
		return "Category updated successfully"

	case "DELETE":
		if len(parts) < 2 {
			return "Error: Missing ID"
		}
		id := parts[1]
		_, err := s.Repo.Delete(id)
		if err != nil {
			return "Error: " + err.Error()
		}
		return "Category deleted successfully"

	default:
		return "Error: Unknown command"
	}
}
