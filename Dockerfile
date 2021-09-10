FROM golang:1.16-alpine

LABEL maintainer="Ram"
WORKDIR /app
COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /habittracker

EXPOSE 8080

CMD ["/habittracker"]
