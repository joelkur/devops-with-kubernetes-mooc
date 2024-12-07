FROM golang:1.23.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o log-output .
ENTRYPOINT [ "./log-output" ]
