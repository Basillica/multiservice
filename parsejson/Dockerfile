## Build
FROM golang:1.17.2-buster AS build

WORKDIR /parsejson
COPY . .

COPY go.mod ./
COPY app.env ./
RUN go mod download

COPY *.go ./
# App api definitions
COPY api/*.go ./api/
# handlers
COPY handlers/*.go ./handlers/
COPY middleware/*.go ./middleware/
COPY types/appenv/*.go ./types/appenv/
COPY types/*.go ./types/

RUN export GIT_COMMIT=$(git rev-list -1 HEAD) && export GIN_MODE=release & \
  go build -ldflags "-X main.GitVersion=$GIT_COMMIT" -o /multiservice_json_parser

## Deploy

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY app.env ./
COPY json/*.json ./json/
COPY --from=build /multiservice_json_parser /multiservice_json_parser
EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/multiservice_json_parser"]