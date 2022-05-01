FROM golang:1.18 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

# just copy main.go for now
# if the app becomes more complex 
# this will need to be changed
COPY go.mod go.sum cmd/main.go ./
COPY pkg/ ./pkg
RUN go mod download && go mod verify

RUN go build -a -o app .


FROM alpine:latest

EXPOSE 8080

WORKDIR /app

COPY --from=builder /build/app .

ENTRYPOINT ./app