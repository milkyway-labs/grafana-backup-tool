build:
	cd orchestrator && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-build:
	docker buildx build --platform linux/amd64 -t milkywaylabs/grafana-backup-tool:v1.0.0 --load .
