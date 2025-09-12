package dto

type SignupreqBody struct {
	Firstname       string `json:"firstname" validate:"required"`
	Lastname        string `json:"lastname" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,alphanum"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,alphanum"`
}
