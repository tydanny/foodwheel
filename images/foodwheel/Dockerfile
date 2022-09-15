FROM golang:1.19 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
COPY pkg/ ./pkg/
COPY cmd/ ./cmd/

RUN go build -a -o server ./cmd/main.go


FROM gcr.io/distroless/static

EXPOSE 50051

WORKDIR /app
COPY --from=builder --chown=nonroot:nonroot /build/server .
USER nonroot

ENTRYPOINT ["/app/server"]
