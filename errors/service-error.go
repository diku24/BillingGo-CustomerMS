package errors

//ServiceError should be return buisness error
type ServiceError struct {
	Message string `json:"message"`
}
