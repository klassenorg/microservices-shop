FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /recommendationservice ./cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /recommendationservice /recommendationservice

EXPOSE 8084

ENTRYPOINT ["/recommendationservice"]