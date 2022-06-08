package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/handlers"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/core/api/routes"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/server/config"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/database/mongodb"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/middlewares"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
)

type Server struct {
	Addr       string
	Port       int
	httpServer http.Server
}

func New(addr string, port int) *Server {
	return &Server{
		Addr: addr,
		Port: port,
	}
}

func (s *Server) Run(cfg config.Config, log logger.ILoggerProvider) error {
	log.Info("server.main.Run", fmt.Sprintf("Server running on port :%d", s.Port))
	log.Info("server.main.Run", fmt.Sprintf("Environment: %s", cfg.Environment))

	//Mongo db conection
	ctx := context.TODO()

	connection := mongodb.New(ctx)
	if connection.Error != nil {
		panic(fmt.Sprintf("error connecting to database: %v", connection.Error))
	} else {
		fmt.Println("DB connected")
	}

	//Start server with config

	handlerDependencies := handlers.Dependencies{Logger: log}

	router := routes.NewRoutes(handlers.NewHandler(handlerDependencies))
	router.Setup()

	s.httpServer = http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Handler:      middlewares.Recovery(router.Client),
		ReadTimeout:  cfg.Application.ReadTimeout * 2,
		WriteTimeout: cfg.Application.WriteTimeout * 2,
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			fmt.Println(err)
			if err == http.ErrServerClosed {
				return
			} else {
				fmt.Println(err)
			}

			process, err := os.FindProcess(os.Getpid())
			if err != nil {
				fmt.Println("couldn't find process to exit")
				os.Exit(1)
			}

			if err := process.Signal(os.Interrupt); err != nil {
				fmt.Println("couldn't find process to exit")
				os.Exit(1)
			}

		}
	}()

	return nil
}
