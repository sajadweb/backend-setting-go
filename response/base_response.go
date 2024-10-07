package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total,omitempty"`
	Error   bool        `json:"error"`
}

func NewBaseResponse(code int, message string, data interface{}, total int, err bool) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Total:   total,
		Error:   err,
	}
}
