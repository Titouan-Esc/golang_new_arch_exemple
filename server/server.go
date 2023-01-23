package server

import (
	"context"
	"net/http"
	"sync"
)

type Connection struct {
	writer     http.ResponseWriter
	flusher    http.Flusher
	requestCtx context.Context
}

type Server struct {
	connections      map[string]*Connection
	connectionsMutex sync.RWMutex
}

func NewServer() *Server {
	server := &Server{
		connections: map[string]*Connection{},
	}
	return server
}

func (s *Server) ServerHTTP(res http.ResponseWriter, req *http.Request) {
	flusher, ok := res.(http.Flusher)
	if !ok {
		http.Error(res, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	requestContext := req.Context()
	s.connectionsMutex.Lock()
	s.connections[req.RemoteAddr] = &Connection{
		writer:     res,
		flusher:    flusher,
		requestCtx: requestContext,
	}
	s.connectionsMutex.Unlock()

	defer func() {
		s.RemoveConnection(req.RemoteAddr)
	}()

	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	<-requestContext.Done()
}

func (s *Server) Send(msg string) {
	s.connectionsMutex.RLock()
	defer s.connectionsMutex.RUnlock()

	msgBytes := []byte("event: message\n\ndata:" + msg + "\n\n")
	for client, connection := range s.connections {
		_, err := connection.writer.Write(msgBytes)
		if err != nil {
			s.RemoveConnection(client)
			return
		}

		connection.flusher.Flush()
	}
}

func (s *Server) RemoveConnection(client string) {
	s.connectionsMutex.Lock()
	defer s.connectionsMutex.Unlock()

	delete(s.connections, client)
}
