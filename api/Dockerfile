FROM golang:1.20.3

# RUN apt-get update \
#     && apt-get install -y --no-install-recommends

# WORKDIR /go/src/

# RUN go mod init github.com/takumi-ishisaka/game-echo-server
# RUN go install github.com/cosmtrek/air@latest
# RUN go mod tidy

# CMD ["air", "-c", ".air.toml"]
# Dockerfile


WORKDIR /go/src
ENV GO111MODULE=on

COPY . .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]