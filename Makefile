
.PHONY: deps
deps:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.18.0
	go get github.com/jstemmer/go-junit-report
	go get github.com/maxcnunes/waitforit
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

.PHONY: openapi
openapi:
	oapi-codegen --package api api/spec.yaml > api/api.gen.go

.PHONY: local
local:
	go run cmd/api/main.go

.PHONY: docker-build
docker-build:
	docker build -t bootstrap-echo-openapi-docker .
