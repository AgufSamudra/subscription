package models

type Login struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type Register struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

// response
type RegisterResponse struct {
	Email string `json:"email"`
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
