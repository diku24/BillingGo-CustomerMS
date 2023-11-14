package errors

//HandlerError should be return buisness error
type ControllerError struct {
	Message string `json:"message"`
}
