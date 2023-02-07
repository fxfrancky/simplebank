# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# Install curl before using it
RUN apk add curl
# Download the migrate file
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
          # sudo mv migrate /usr/bin

# Run stage (to have a small binary result)
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
# Copy the migrated file
COPY --from=builder /app/migrate ./migrate
# COPY Migration Files
COPY db/migration ./migration
# Copy the app.en file
COPY app.env .
COPY start.sh .
COPY wait-for.sh .


EXPOSE 8080
CMD [ "/app/main" ]
# Execute start.sh and use CMD as a parameter '/app/main'
ENTRYPOINT [ "/app/start.sh" ]