package types

type Config struct {
	ListenAddress string
	BasePath      string
	EnableSwagger bool
	ChunkSize     int64
}

type HttpResponse struct {
	Error   bool      `json:"error"`
	Message string    `json:"message"`
	Details *[]string `json:"details"`
}
