#Build stage
FROM golang:1.22.9-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Runs stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
