package request

type RegisterRequest struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}
