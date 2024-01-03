.PHONY: build-http
build-http: build-docs build-api

.PHONY: build-docs
build-docs:
	swag init --parseDependency --parseInternal -g http.go -d "./internal/handlers/http" -o "./internal/handlers/http/docs"

.PHONY: build-api
build-api:
	GOOS=linux GOARCH=amd64 go build -o ./out/http ./cmd/http/main.go

.PHONY: build-lambda
build-lambda:
	GOOS=linux GOARCH=arm64 go build -o ./out/bootstrap ./cmd/lambda/main.go

.PHONY: build-auth4
build-auth4:
	GOOS=linux GOARCH=arm64 go build -o ./out/bootstrap ./cmd/auth4/main.go