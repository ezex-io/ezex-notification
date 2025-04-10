# Build stage
FROM golang:1.24-alpine AS builder

RUN apk --no-cache add make

WORKDIR /app
COPY . .

RUN make release

FROM alpine:latest

RUN mkdir /etc/notification
COPY --from=builder /app/build/ezex-notification /usr/bin/ezex-notification

EXPOSE 50051

ENTRYPOINT ["/usr/bin/ezex-notification"]
