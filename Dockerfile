FROM golang:1.19 as build

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build ./cmd/main.go

# Run
CMD ["./main"]

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app /app
EXPOSE 8080
CMD ["./app/main"]