package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,gte=6,lte=100"`
	Password string `json:"password" binding:"required,gte=6,lte=100"`
}
