package structs

type DefaultResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorReponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
