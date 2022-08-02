package respons

type respons struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIRespons(message string, code int, status string, data interface{}) respons {
	meta := meta{}
	meta.Message = message
	meta.Code = code
	meta.Status = status

	respons := respons{}
	respons.Meta = meta
	respons.Data = data

	return respons
}
