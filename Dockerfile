FROM golang:alpine3.15

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /app-runner

EXPOSE 3000

CMD [ "/app-runner" ]
