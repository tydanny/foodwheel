FROM golang:1.24 AS build-stage

ARG LDFLAGS

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="${LDFLAGS}" -o /foodwheel

FROM gcr.io/distroless/static

WORKDIR /

EXPOSE 50051

USER nonroot

COPY --from=build-stage --chown=nonroot:nonroot /foodwheel /foodwheel

ENTRYPOINT ["/foodwheel"]
