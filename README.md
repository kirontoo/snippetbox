# SnippetBox

This is a learning project from the book "Let's Go" by Alex Edwards.

## Run Server Locally
```bash
go run ./cmd/web
```

## Genereate TLS certificate
```go
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
