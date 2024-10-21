package request

type LoginUserReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"pwd" binding:"required,min=6"`
}

type RegisterUserReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"pwd" binding:"required,min=6"`
}
