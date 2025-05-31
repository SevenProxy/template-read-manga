FROM golang:alpine AS Builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=Builder /app/main .

EXPOSE 3001

CMD ["./main"]