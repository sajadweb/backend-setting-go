package common

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	// "net"
)

type Response struct {
	Disposed bool        `json:"isDisposed"`
	ID       string      `json:"id"`
	Response     interface{} `json:"response"`
	Err      interface{} `json:"err"` // Can be nil or other types
}
type DataResponse struct {
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
	ErrMsg string      `json:"message"`
	Error  interface{} `json:"error"`
}

// func to convert the struct to a JSON string
func ConvertToString(response Response) (string, error) {
	// Use json.Marshal to convert struct to JSON string
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", err
	}
	// Get the length of jsonBytes
	jsonLength := len(jsonBytes)
	// Return the JSON string
	return fmt.Sprintf("%v#%v", jsonLength, string(jsonBytes)), nil
}
func SendResponse(request *TcpRequest, code int, data interface{}, message string, err bool) {
	responseData := DataResponse{
		Data:   data,
		Code:   code,
		ErrMsg: message,
		Error:  err,
	}
	response := Response{
		ID:       request.ID,
		Disposed: true,
		Response:  responseData,
		Err:      nil,
	}
	res,_ := ConvertToString(response)
	fmt.Println(res)
	_, writeErr := request.conn.Write([]byte(res))
	if writeErr != nil {
		return
	}
}
