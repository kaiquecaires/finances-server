FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o finances .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/finances .
COPY .env .
EXPOSE 8080
CMD ["./finances"]
