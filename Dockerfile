FROM golang:1.21

WORKDIR /app
COPY . .
RUN go build -o main
CMD ["air", "-c", ".air.toml"]