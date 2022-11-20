package types

type Config struct {
	ListenAddress string
	BasePath      string
}

type HttpResponse struct {
	Error   bool      `json:"error"`
	Message string    `json:"message"`
	Details *[]string `json:"details"`
}
