# Build container
FROM golang:1.19-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /todo

# Retrieve application dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY .. ./
RUN go build -v -o monolith ./cmd/todo-api

# Run the Go Binary in Alpine.
FROM alpine:latest AS runtime
COPY --from=builder /todo/docker/wait-for /todo/wait-for
RUN chmod +x /todo/wait-for
COPY --from=builder /todo/monolith /todo/monolith
COPY --from=builder /todo/environments /todo/environments
WORKDIR /todo
CMD ["/monolith"]
