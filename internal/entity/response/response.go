package response

type FieldError struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

type ValidationErrors struct {
	Errors []FieldError `json:"errors"`
}

type HttpResponse struct {
	Message string           `json:"message,omitempty"`
	Status  string           `json:"status,omitempty"`
	Data    string           `json:"data,omitempty"`
	Errors  ValidationErrors `json:"errors,omitempty"`
}
