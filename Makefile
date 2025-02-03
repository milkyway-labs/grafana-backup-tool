build:
	cd orchestrator && CGO_ENABLED=0 GOOS=linux go build -o bin/orchestrator
	docker build -t grafana-backup -f Dockerfile .
