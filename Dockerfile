# ── Build stage ──────────────────────────────────────────────
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w -trimpath" \
    -o cv-terminal ./cmd/server

# ── Run stage ────────────────────────────────────────────────
FROM alpine:3.21

RUN addgroup -S app && adduser -S -u 10001 -G app app

WORKDIR /app
COPY --from=builder /app/cv-terminal .

# Dossier pour la clé SSH hôte (monté en volume sur le VPS)
RUN mkdir -p .ssh && chown -R app:app /app

USER app
EXPOSE 22

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
    CMD nc -z localhost 22 || exit 1

ENV SSH_HOST=0.0.0.0
ENV SSH_PORT=22
ENV SSH_HOST_KEY_PATH=.ssh/id_ed25519

ENTRYPOINT ["./cv-terminal"]
