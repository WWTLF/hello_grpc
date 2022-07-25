# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/hello_client && touch jwt_signing_key.pem

# final stage
FROM alpine:latest
COPY nsswitch.conf /etc/nsswitch.conf
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/hello_client /app/
COPY --from=builder /app/jwt_signing_key.pem /etc/ssl/private/jwt_signing_key.pem
EXPOSE 50031
CMD ["/app/hello_client"]%  
