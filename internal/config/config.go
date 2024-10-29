package config

import "sync/atomic"

// Hold stateful, in-memory data
// Use atomic std lib to prevent data racing casued by concurrent requests
type ApiConfig struct {
	FileServerHits atomic.Int32
}
