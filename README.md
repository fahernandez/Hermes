# Introduction
This repo contains Hermes, an SMS messaging application.

# Getting started

## Requirements
- [Docker](https://docs.docker.com/engine/installation/)
- [Docker compose](https://docs.docker.com/compose/install/)

## Installation

#### Hermes standalone (for app developers)
To use Hermes as a standalone app for development, simply run the docker compose config in this repo: `docker-compose up -d`

##### Fetch dependencies
To fetch any missing dependencies run the following command:
```bash
docker run --rm -v "$PWD":/go/src/github.com/fahernandez/hermes -w /go/src/github.com/fahernandez/hermes huli/golang sh -c "go get . && govendor init && govendor add +external"
```

##### Build
To compile Hermes run the following command:
```bash
docker run --rm -v "$PWD":/go/src/github.com/fahernandez/hermes -w /go/src/github.com/fahernandez/hermes huli/golang go build -v
```

##### Run tests
Before running the Hermes tests, make sure to have the services it depends on running, you can start them with `docker-compose up -d`. Then use the following command to run the Hermes tests:
```bash
docker run --rm  --net=hermes_default -v "$PWD":/go/src/github.com/fahernandez/hermes -w /go/src/github.com/fahernandez/hermes huli/golang go test -cover -v ./...
```