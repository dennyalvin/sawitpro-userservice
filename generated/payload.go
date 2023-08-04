package generated

type ProfileUpdateParams struct {
	Id       int     `json:"id"`
	Phone    *string `json:"phone"`
	FullName *string `json:"full_name"`
}

type SignupParams struct {
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type LoginParams struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
