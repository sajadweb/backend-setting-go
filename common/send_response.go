package common

import (
	// "encoding/json"
	"fmt"
	"net"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
	Error   bool        `json:"error"`
}

func SendResponse(conn net.Conn, code int, message string, data interface{}, total int, error bool) {
	response := BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Total:   total,
		Error:   error,
	}
	fmt.Println(response)
	// response := models.NewBaseResponse(code, message, data, total, error)
	// respJSON, err := json.Marshal(response)
	// if err != nil {
	// 	// Handle marshaling error, you might want to send an error response back here
	// 	conn.Write([]byte(`{"code": 500, "message": "Internal Server Error", "error": true}`))
	// 	return
	// }
	// conn.Write(response)
	// conn.Write([]byte("\n"))
	_ , err := conn.Write([]byte(`{"code": 500, "message": "Internal Server Error123", "error": true}`))
	if err != nil {
		fmt.Printf("hasserrr in response %v",err)
	}	
}
