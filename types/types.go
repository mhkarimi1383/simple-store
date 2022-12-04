package types

// used to get Store configurations
type Config struct {
	ListenAddress string
	BasePath      string
	EnableSwagger bool
	ChunkSize     int64
}

// the base type for returning data to client
type HttpResponse struct {
	Error   bool      `json:"error"`
	Message string    `json:"message"`
	Details *[]string `json:"details"`
}
