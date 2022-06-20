FROM golang:1.18-alpine 
EXPOSE 8080
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o application -v cmd/main.go
CMD ["/app/application"]
