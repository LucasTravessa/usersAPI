package response

type UserResponse struct {
	ID    string `json:"id"` // uuid
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}