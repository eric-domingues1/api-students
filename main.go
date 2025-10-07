package main

import (
	"github.com/eric-domingues1/api-students/api"
	"github.com/rs/zerolog/log"
)

func main() {
	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server")
	}
}
