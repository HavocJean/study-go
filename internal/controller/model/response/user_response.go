package response

type UserResponse struct {
	ID    *uint  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
