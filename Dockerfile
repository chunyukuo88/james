FROM 1.19-alpine3.16

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 3000

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]