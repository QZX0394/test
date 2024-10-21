package response

type LoginResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"msg"`
	Data    LoginData `json:"data"`
}

type RegisterResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"msg"`
	Data    RegisterData `json:"data"`
}

type LoginData struct {
	Token string `json:"token"`
}

type RegisterData struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}
