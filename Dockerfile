# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage (to have a small binary result)
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
# Copy the app.en file
COPY app.env .


EXPOSE 8080
CMD [ "/app/main" ]