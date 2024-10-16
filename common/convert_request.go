package common

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// Define the struct for the JSON part
type TcpRequest struct {
	Pattern string                 `json:"pattern"`
	Data    map[string]interface{} `json:"data"` // Use map for flexible data
	ID      string                 `json:"id"`
	conn    net.Conn
}
 
// Function to convert the string to the struct
func MakeTcpRequest(conn net.Conn, message string) (*TcpRequest, error) {
	// Split the message on the #
	parts := strings.SplitN(message, "#", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format")
	}

	// Parse the JSON part
	var request TcpRequest
	err := json.Unmarshal([]byte(parts[1]), &request)
	if err != nil {
		return nil, err
	}
	request.conn = conn

	// Return the populated struct
	return &request, nil
}
