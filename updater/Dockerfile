FROM golang:1.22 AS build
LABEL maintainer="oyyblin <4529189+oyyblin@users.noreply.github.com>"
WORKDIR /src
COPY . .

RUN apt update
RUN apt install -y build-essential

COPY ./go.mod ./go.sum ./

RUN GIT_TERMINAL_PROMPT=1 \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o /bin/app ./cmd/main.go

FROM gcr.io/distroless/cc-debian12
COPY --from=build /bin/app /

# Metrics port
EXPOSE 4014

ENTRYPOINT ["/app"]
