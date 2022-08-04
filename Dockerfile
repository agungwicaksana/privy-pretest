FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go


FROM alpine:3

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT [ "/app/main" ]