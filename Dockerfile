FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# just copy main.go for now
# if the app becomes more complex 
# this will need to be changed
COPY main.go .

EXPOSE 8080

CMD go run .