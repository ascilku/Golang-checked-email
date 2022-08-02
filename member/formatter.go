package member

type formatter struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Password string `json:"-" `
	Token    string `json:"token"`
}

func Formatter(member Member, token string) formatter {
	formatter := formatter{}
	formatter.Id = member.Id
	formatter.Nama = member.Nama
	formatter.Password = member.Password
	formatter.Token = token

	return formatter
}
