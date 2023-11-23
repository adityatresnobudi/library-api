package dto

type UserPayload struct {
	ID    int    `json:"id"`
	Name  string `json:"user_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Password string `json:"user_password" binding:"required"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"user_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
