# Builder
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

# Staging
FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/
RUN apk add --no-cache gcompat

RUN chmod +x ./main

EXPOSE 8080

CMD [ "./main" ]