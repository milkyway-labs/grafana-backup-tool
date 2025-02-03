package app

import (
	"encoding/json"
	"net/http"
	"orchestrator/logger"
	"os"
	"os/exec"
)

func runCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)

    logger.Info("Running command: " + cmd.String())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func successResponse(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)
	json.NewEncoder(resp).Encode(Response{Status: 0, Error: ""})
}
func jsonResponse(resp http.ResponseWriter, data interface{}) {
    resp.Header().Set("Content-Type", "application/json")
    resp.WriteHeader(200)
    json.NewEncoder(resp).Encode(data)
}
func failResponse(resp http.ResponseWriter, err error) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(500)
	json.NewEncoder(resp).Encode(Response{Status: -1, Error: err.Error()})
}
