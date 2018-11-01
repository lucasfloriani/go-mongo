package helper

type Response struct {
	Error    string      `json:"error"`
	Response interface{} `json:"response"`
}

func NewErrorResponse(err error) Response {
	return Response{Error: err.Error()}
}

func NewSuccessResponse(r interface{}) Response {
	return Response{Response: r}
}
