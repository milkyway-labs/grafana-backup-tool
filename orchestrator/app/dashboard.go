package app

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"orchestrator/logger"
	"os"
	"strings"
)

type DashboardsInfo map[string]string

// Read a file in the tar.gz archive
func readFile(zipFilePath string, targetFilePath string) (*string, error) {
	file, err := os.Open(zipFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	var fileContent bytes.Buffer
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if header.Name == targetFilePath {
            logger.Info("Found file: " + header.Name)

			if _, err := io.Copy(&fileContent, tr); err != nil {
				return nil, err
			}

            logger.Info(fmt.Sprintf("File %s extracted successfully", header.Name))

			fileStr := fileContent.String()
			return &fileStr, nil
		}
	}

	return nil, errors.New("File not found")
}

func getDashboardsInfo(archive string) (DashboardsInfo, error) {
	archiveName := strings.Split(archive, ".")[0]
	zipFilePath := fmt.Sprintf("./_OUTPUT_/%s", archive)
	targetFilePath := fmt.Sprintf("_OUTPUT_/dashboards/%s/dashboards_%s.txt", archiveName, archiveName)

	file, err := readFile(zipFilePath, targetFilePath)
    if err != nil {
        return nil, err
    }

    dashboardsInfo := make(DashboardsInfo)
	info := strings.Split(*file, "\n")
	// File has an empty line at the end
	for _, dashboard := range info[0 : len(info)-1] {
		dashboardInfo := strings.Split(dashboard, "\t")
		id := strings.Split(dashboardInfo[0], "/")[1]
		dashboardsInfo[dashboardInfo[1]] = id
	}

	return dashboardsInfo, nil
}

func getDashboard(archive string, title string) (*DashboardResponse, error) {
    dashboardsInfo, err := getDashboardsInfo(archive)
    if err != nil {
        return nil, err
    }

    dashboardId, exists := dashboardsInfo[title]
    if !exists {
        msg := fmt.Sprintf("Dashboard with title '%s' not found", title)
        return nil, errors.New(msg)
    }

	archiveName := strings.Split(archive, ".")[0]
	zipFilePath := fmt.Sprintf("./_OUTPUT_/%s", archive)
    targetFilePath := fmt.Sprintf("_OUTPUT_/dashboards/%s/%s.dashboard", archiveName, dashboardId)

	file, err := readFile(zipFilePath, targetFilePath)
    if err != nil {
        return nil, err
    }

	var dashbaordFile map[string]interface{}
	if err := json.Unmarshal([]byte(*file), &dashbaordFile); err != nil {
        return nil, err
	}

    dashboardResponse := DashboardResponse{}
	dashboard, ok := dashbaordFile["dashboard"]
	if !ok {
        return nil, errors.New("Dashboard key not found")
	}

	dashboardBytes, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
        return nil, err
	}

    if err := json.Unmarshal(dashboardBytes, &dashboardResponse); err != nil {
        return nil, err
    }
    
    return &dashboardResponse, nil
}
