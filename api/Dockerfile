FROM golang:1.20.3

RUN apt-get update \
    && apt-get install -y --no-install-recommends

WORKDIR /go/src
ENV GO111MODULE=on

COPY ./src/go.mod ./src/go.sum ./
RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.43.0

CMD ["air", "-c", ".air.toml"]