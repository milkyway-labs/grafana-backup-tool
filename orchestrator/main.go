package main

import (
	"errors"
	"net/http"
	"orchestrator/app"
	"orchestrator/logger"
	"os"
)

func main() {
	err := app.Run()

	if errors.Is(err, http.ErrServerClosed) {
		logger.Info("Server closed")
	} else if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
