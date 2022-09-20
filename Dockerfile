FROM golang:1.18-alpine:latest as builder
WORKDIR /statistics
COPY . .
RUN go build -o main main.go


FROM alpine:latest
WORKDIR /statistics
COPY --from=builder /receiver/main . 



CMD ["statistics/main"]