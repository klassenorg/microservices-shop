FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /pricingservice ./cmd/main/main.go

FROM scratch

WORKDIR /

COPY --from=build /pricingservice /pricingservice

EXPOSE 8083

ENTRYPOINT ["/pricingservice"]