package model

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}
