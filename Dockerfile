FROM golang:1.18-alpine3.14 as builder
WORKDIR /statistics
COPY . .
RUN go build -o main main.go


FROM alpine:3.14
WORKDIR /statistics
COPY --from=builder /receiver/main . 



CMD ["statistics/main"]
