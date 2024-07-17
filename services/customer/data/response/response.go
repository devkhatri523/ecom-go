package response

type Response struct {
	Code   int
	Status string
	Data   interface{}
}

type ErrorResponse struct {
	Code    int
	Message string
}
