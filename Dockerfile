FROM golang:1.22.0-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o lb-build

FROM golang:1.22.0-alpine AS prod
WORKDIR /app
COPY --from=build /app/lb-build ./
COPY --from=build /app/config.json ./

EXPOSE 2703
CMD ["./lb-build", "config.json"]