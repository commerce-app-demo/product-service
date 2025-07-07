FROM golang:1.24.4-alpine AS build

WORKDIR /go/src/product-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/product-service cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/product-service /

EXPOSE 50051
CMD [ "/product-service" ]

