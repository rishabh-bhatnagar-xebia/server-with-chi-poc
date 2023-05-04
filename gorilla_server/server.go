package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	Config Config
}

type Config struct {
	Port string
}

func (s *Server) Start() error {
	router := s.routes()
	fmt.Println("Running a server on port:", s.Config.Port)
	return http.ListenAndServe(":"+s.Config.Port, router)
}
