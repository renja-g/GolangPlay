package request

type SignupUserRequest struct {
	Email string `validate:"required,min=1,max=200" json:"email"`
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}
