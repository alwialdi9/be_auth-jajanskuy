FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN sed -i 's/http\:\/\/dl-cdn.alpinelinux.org/https\:\/\/alpine.global.ssl.fastly.net/g' /etc/apk/repositories
RUN apk --no-cache add curl git ca-certificates tzdata && \
    ls -la /usr/share/zoneinfo/Asia/Jakarta

COPY . .
RUN go mod tidy

RUN go build -o app ./main.go

FROM alpine:latest

# Set up non-root user for security
RUN adduser -D appuser

# Set working directory and copy binary from build stage
WORKDIR /app
COPY --from=builder /app .

# Change ownership and permissions for the app user
RUN chown -R appuser /app
USER appuser

# Expose port and set the entrypoint
EXPOSE 3000
# COPY timezone data and certs
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./app"]