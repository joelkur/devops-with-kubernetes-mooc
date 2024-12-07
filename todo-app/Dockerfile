FROM golang:1.23.3-alpine
WORKDIR /app
COPY . .
RUN go build -o todo-app .
ENTRYPOINT [ "./todo-app" ]
