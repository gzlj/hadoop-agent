package module

type BusinessResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func BuildResponse(code int, message string, data interface{}) (response BusinessResponse){
	response = BusinessResponse{
		Code: code,
		Message: message,
		Data: data,
	}
	return
}
