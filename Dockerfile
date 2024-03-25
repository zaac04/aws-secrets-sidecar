FROM golang:alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o SecretSideCar .
FROM alpine:latest
WORKDIR /app
COPY --from=build /app .
CMD [ "./SecretSideCar" ]
