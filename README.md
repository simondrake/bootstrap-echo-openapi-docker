# bootstrap-echo-openapi-docker

Bootstrap repo for echo, openAPI and Docker. To be used for example/reference.

# Purpose

The purpose of this repo is to provide a working example of the following:

* Updating `api/spec.yaml` and running `make openapi` will generate `api/api.gen.go`
* The code in `cmd/api/main.go` spins up a HTTP server, using the Swagger definition in `api/api.gen.go` and the Echo framework
* The `Server` interface in `internal/server/v*/server.go` has to satisy the Interface in `api/api.gen.go`
* `Dockerfile` builds the code into a Docker container
* `docker-compose.yaml` shows how to use the Docker container


# Developmet

## Publish to docker

```
docker tag bootstrap-echo-openapi-docker:latest omisnomis/bootstrap-echo-openapi-docker:latest

docker push omisnomis/bootstrap-echo-openapi-docker:latest
```
