package response

type Response struct {
	Success bool              `json:"success"`
	Data    map[string]string `json:"data"`
	Error   map[string]string `json:"error"`
}

func Success(data map[string]string) Response {
	return Response{
		Success: true,
		Data:    data,
		Error:   map[string]string{},
	}
}

func Error(msg string) Response {
	return Response{
		Success: false,
		Data:    map[string]string{},
		Error:   map[string]string{"msg": msg},
	}
}
