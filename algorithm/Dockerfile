## Build
FROM golang:1.17.2-buster AS build

WORKDIR /algorithm
COPY . .

COPY *.go ./

RUN export GIT_COMMIT=$(git rev-list -1 HEAD) && export GIN_MODE=release & \
  go build -ldflags "-X main.GitVersion=$GIT_COMMIT" -o /multiservice_algorithm

## Deployment
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY app.env ./
COPY --from=build /multiservice_algorithm /multiservice_algorithm
EXPOSE 8082

USER nonroot:nonroot

ENTRYPOINT ["/multiservice_algorithm"]