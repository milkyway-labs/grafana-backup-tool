package app

import (
	"fmt"
	"net/http"
	"orchestrator/logger"
)

func mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/backup", func(resp http.ResponseWriter, req *http.Request) {
		command := "grafana-backup save"
		if err := runCommand(command); err != nil {
			logger.Error(err)
			failResponse(resp, err)

			return
		}
		successResponse(resp)
	})

	mux.HandleFunc("/restore/all", func(resp http.ResponseWriter, req *http.Request) {
		archive := req.URL.Query().Get("archive")

		if archive == "" {
			err := fmt.Errorf("archive parameter is required")
			logger.Error(err)
			failResponse(resp, err)

			return
		}

        exist, err := isFileExist(archive)
        if err != nil {
            logger.Error(err)
            failResponse(resp, err)

            return
        }

		command := fmt.Sprintf("grafana-backup restore %s", archive)
        if !exist {
            logger.Info("Archive not found in aws s3")
            logger.Info("Try restore in local")

            command = fmt.Sprintf("unset AWS_S3_BUCKET_NAME && grafana-backup restore %s/%s", "_OUTPUT_", archive)
        }

		if err := runCommand(command); err != nil {
			logger.Error(err)
			failResponse(resp, err)

			return
		}

		successResponse(resp)
	})

	mux.HandleFunc("/restore/dashboard", func(resp http.ResponseWriter, req *http.Request) {
		archive := req.URL.Query().Get("archive")
		title := req.URL.Query().Get("title")

		if archive == "" || title == "" {
			err := fmt.Errorf("archive, title parameter is required")
			logger.Error(err)
			failResponse(resp, err)

			return
		}

		dashboard, err := getDashboard(archive, title)
		if err != nil {
			logger.Error(err)
			failResponse(resp, err)

			return
		}

        jsonResponse(resp, dashboard)
	})

	return mux
}
