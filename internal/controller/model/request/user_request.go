package request

type UserRequest struct {
	Email    string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}
