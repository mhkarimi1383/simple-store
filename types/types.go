// Package types all of the types used in project are here
package types

// Config used to get Store configurations
type Config struct {
	ListenAddress string
	BasePath      string
	EnableSwagger bool
	ChunkSize     int64
}

// HTTPResponse the base type for returning data to client
type HTTPResponse struct {
	Error   bool      `json:"error"`
	Message string    `json:"message"`
	Details *[]string `json:"details"`
}
