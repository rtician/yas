package types

type UserRegistration struct {
	Name            string `json:"name" validate:"required" trans:"name"`
	Email           string `json:"email" validate:"required,email" trans:"email"`
	Password        string `json:"password" validate:"required" trans:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required" trans:"confirm_password"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" trans:"email"`
	Password string `json:"password" validate:"required" trans:"password"`
}
