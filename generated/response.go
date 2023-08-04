package generated

// APIResponse General Restful Response wrapper standard
type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseError General Restful Response for Error standard
type ResponseError struct {
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors"`
}
type ErrorDetail struct {
	Title   string `json:"field"`
	Message string `json:"message"`
}

// ProfileResponse User Response for exposed to client
type ProfileResponse struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

// LoginResponse Response struct for Login
type LoginResponse struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
