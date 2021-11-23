FROM golang:1.17-alpine3.14

WORKDIR /kelas-e-app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o mainfile

EXPOSE 9900

CMD ["mainfile"]