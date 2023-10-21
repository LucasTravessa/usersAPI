package request

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,containsany=+-=!@#$%^&*()_+,min=6,max=50"`
}
