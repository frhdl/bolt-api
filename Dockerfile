FROM golang:1.16 AS build

WORKDIR /app
COPY . .

RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend main.go 

FROM alpine:3.12

COPY --from=build /app/backend .
COPY --from=build /app/migrations/schema.sql ./migrations/schema.sql
CMD ["./backend", "web"]