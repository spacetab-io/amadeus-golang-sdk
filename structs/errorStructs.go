package structs

// ErrorReply Error Reply structure
type ErrorReply struct {
	Status    string `json:"status,omitempty"`
	ErrorCode string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}
