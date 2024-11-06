FROM  golang:1.23 as builder
WORKDIR /app
COPY go.* ./ 
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a \
  -ldflags="-w -s -X main.version=$(git describe --tags --always --dirty)" \
  -o ./app ./cmd/main.go

FROM scratch as runtime
COPY --from=builder /app/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./app"]