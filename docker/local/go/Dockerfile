FROM golang:1.20.3

RUN apt-get update \
    && apt-get install -y --no-install-recommends

WORKDIR /var/www
RUN go mod init github.com/takumi-ishisaka/game-echo-server
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]