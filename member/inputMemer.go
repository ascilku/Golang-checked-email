package member

type InputMember struct {
	Nama     string `json:"nama" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginMember struct {
	Nama     string `json:"nama" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Nama string `json:"nama" binding:"required,email"`
}
