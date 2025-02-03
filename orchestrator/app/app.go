package app

import (
	"net/http"
	"orchestrator/logger"
	"os"
)

func Run() error {
    port, exist := os.LookupEnv("PORT")

    if !exist {
        port = "9111"
    }

    logger.Info("Server is running on port :" + port)
    err := http.ListenAndServe(":"+port, mux())

    return err
}
