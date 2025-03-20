run:
	APP_ENV=local go run cmd/multimodal-search/main.go
build:
	go build -o multimodal-search cmd/multimodal-search/main.go
docker-build:
	docker build --build-arg APP_ENV=dev -t multimodal-search .
run-w-compose:
	docker compose up -d --build
run-tests:
	go test -count=1 ./...
lint:
	gofmt -w .
	goimports -w .
	golangci-lint run -v