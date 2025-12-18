package entity

type UserEntity struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Age      int8   `bson:"age"`
}
