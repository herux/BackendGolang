package server

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herux/indegooweather/config"
)

type Server struct {
	srv  *fiber.App
	conf *config.ServerConfig
}

type AddRoute = func(fiber.Router)

func SetupService(conf *config.ServerConfig, addRoute AddRoute) *Server {
	srv := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(conf.ReadTimeout),
	})

	// srv.Use() // attach middleware here, eg: CORS
	addRoute(srv)
	return &Server{
		srv:  srv,
		conf: conf,
	}
}

func (s *Server) Run() {
	s.listen()
}

func (s *Server) listen() {
	url := fmt.Sprintf(":%d", s.conf.Port)

	if err := s.srv.Listen(url); err != nil {
		log.Fatalf("server exited from listen() %s", err)
	}
}
