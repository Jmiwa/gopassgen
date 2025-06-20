FROM golang:1.20-bullseye AS builder
WORKDIR /app
ARG CGO_ENABLED=0
COPY . .
RUN go build -o gopassgen ./cmd/gopassgen/main.go

FROM scratch
ARG VERSION=1.0.0
LABEL org.opencontainers.image.source=https://github.com/Jmiwa/gopassgen \
      org.opencontainers.image.version=${VERSION} \
      org.opencontainers.image.title=gopassgen \
      org.opencontainers.image.description="A password generator CLI tool written in Go."

COPY --from=builder /app/gopassgen /opt/gopassgen/gopassgen
COPY --from=golang:1.20 /etc/passwd /etc/passwd
COPY --from=golang:1.20 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /workdir
USER nobody
ENTRYPOINT ["/opt/gopassgen/gopassgen"]
