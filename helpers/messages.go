package helpers

var (
	ErrorMessageBadRequest = "Bad Request"

	BadRequesetMessage  = map[string]interface{}{"error": "Bad Request"}
	SuccessMessage      = map[string]interface{}{"success": true}
	UnauthorizedMessage = map[string]interface{}{"error": "Unauthorized! Please login first"}
)

// Message returns api call
type Message struct {
	Error   error       `json:"error,omitempty"`
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessDataMessage(data interface{}) map[string]interface{} {
	return map[string]interface{}{"data": data}
}
