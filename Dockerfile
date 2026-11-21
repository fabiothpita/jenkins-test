FROM golang:1.25-alpine AS base
WORKDIR /app
COPY go.mod ./
COPY *.go ./

FROM base AS test
RUN go test -v ./...

FROM base AS builder
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8081
CMD ["./server"]