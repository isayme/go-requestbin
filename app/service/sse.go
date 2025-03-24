package service

import (
	"github.com/r3labs/sse/v2"
)

func NewSseServer() *sse.Server {
	server := sse.New()       // create SSE broadcaster server
	server.AutoReplay = false // do not replay messages for each new subscriber that connects
	_ = server.CreateStream("requests")

	return server
}
