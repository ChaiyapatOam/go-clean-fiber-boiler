FROM golang:1.20-alpine

WORKDIR /usr/src/app

RUN apk add tzdata
ENV TZ Asia/Bangkok

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o ./bin/go-fiber
EXPOSE 5000
CMD [ "./bin/go-fiber"]