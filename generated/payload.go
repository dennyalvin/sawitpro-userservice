package generated

type ProfileUpdateParams struct {
	Id       int     `json:"id"`
	Phone    *string `json:"phone" validate:"phone"`
	FullName *string `json:"full_name" validate:"min=3,max=60"`
}

type SignupParams struct {
	Phone    string `json:"phone" validate:"required,phone"`
	FullName string `json:"full_name" validate:"required,min=3,max=60"`
	Password string `json:"password" validate:"required,min=6,max=64,password"`
}

type LoginParams struct {
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}
