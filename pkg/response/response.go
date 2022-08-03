package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(status, message string, data interface{}) (response Response) {
	response = Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return
}

func CreateErrorResponse(status, message string) (response Response) {
	if message == "" {
		message = MESSAGE_ERROR
	}
	response = Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
	return
}
