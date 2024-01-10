package response

type Response struct {
	Status string `json: "status"`
	Error  string `json: "error, omitempty"`
}

const (
	StatusOk  = "OK"
	StatusErr = "Error"
)

func Ok() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(message string) Response {
	return Response{
		Status: StatusErr,
		Error:  message,
	}
}
