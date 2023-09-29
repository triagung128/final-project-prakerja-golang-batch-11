package response

type Auth struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Token       string `json:"token"`
}
