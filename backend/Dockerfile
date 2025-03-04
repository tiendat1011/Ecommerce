FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 8080

CMD [ "./main" ]