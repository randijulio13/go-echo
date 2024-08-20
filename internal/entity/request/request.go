package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required" name:"name"`
	Username string `json:"username" validate:"required" name:"username"`
	Password string `json:"password" validate:"required" name:"password"`
	Email    string `json:"email,omitempty" validate:"required,email" name:"email"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
