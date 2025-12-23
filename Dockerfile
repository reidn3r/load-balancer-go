FROM golang:1.22.0-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o lb-build

FROM golang:1.22.0-alpine AS prod
WORKDIR /app
COPY --from=build /app/lb-build ./

EXPOSE 8080
RUN ["./lb-build"]