package forms

type LoginForm struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type LogoutForm struct {
	Email string `json:"email" form:"email" query:"email"`
	Token string `json:"token" form:"token" query:"token"`
}
