FROM golang:1.20

WORKDIR /app


ENV GIN_MODE="release"

ENV DbHostName="db"
ENV DbPort="5432"
ENV DbName="postgres"
ENV DbUser="postgres"
ENV DbPassword="admin"
ENV TmpFolder = "tmp"

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main

CMD ["./main"]