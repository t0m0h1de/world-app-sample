FROM docker.io/golang:1.23 as builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/app

FROM docker.io/alpine:latest

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
