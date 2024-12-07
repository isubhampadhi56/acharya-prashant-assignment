package main

import (
	"net/http"
	"strconv"

	"github.com/api-assignment/pkg/config"
	router "github.com/api-assignment/pkg/routes"
	"github.com/api-assignment/pkg/utils/logger"
)

func main() {
	router := router.MainRouter()
	port := strconv.Itoa(config.GetConfig().GetAppPort())
	log := logger.InitializeAppLogger()
	log.Info("Starting API Server on Port ", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("unable to start server on port %d ", port, err)
	}
}
