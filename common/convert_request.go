package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Define the struct for the JSON part
type MsgData struct {
	Pattern string                 `json:"pattern"`
	Data    map[string]interface{} `json:"data"` // Use map for flexible data
	ID      string                 `json:"id"`
}

// Function to convert the string to the struct
func Convert(msg string) (*MsgData, error) {
	// Step 1: Split the message on the #
	parts := strings.SplitN(msg, "#", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format")
	}

	// Step 2: Parse the JSON part
	var msgData MsgData
	err := json.Unmarshal([]byte(parts[1]), &msgData)
	if err != nil {
		return nil, err
	}

	// Return the populated struct
	return &msgData, nil
}
type Response struct {
	Pattern bool                 `json:"isDisposed"`
	ID     string                 `json:"id"`
	Status string                 `json:"status"`
	Data   interface{}            `json:"response"`
	Err    interface{}            `json:"err"` // Can be nil or other types
}

//func to convert the struct to a JSON string
func ConvertToString(Pattern string,id string, data interface{}) (string, error) {
	response := Response{
		ID:    id,	
		Pattern: true,
		Status: "success",
		Data: data,
		Err: nil,
	}

	// Use json.Marshal to convert struct to JSON string
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	// Get the length of jsonBytes
	jsonLength := len(jsonBytes)
	fmt.Print(jsonLength)
	// Return the JSON string
	return fmt.Sprintf("%v#%v", jsonLength, string(jsonBytes)), nil
}