package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"` // エラー時のみ使用
	Data    interface{} `json:"data,omitempty"`    // 成功時のみ使用
}

func Success(data interface{}) Response {
	return Response{
		Status: "ok",
		Data:   data,
	}
}

func Error(message string) Response {
	return Response{
		Status:  "ng",
		Message: message,
	}
}
