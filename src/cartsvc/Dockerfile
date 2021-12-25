FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /cartservice ./cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /cartservice /cartservice

EXPOSE 8082

ENTRYPOINT ["/cartservice"]