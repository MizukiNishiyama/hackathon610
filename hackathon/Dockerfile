# ベースイメージを指定する
FROM golang:1.18 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o testapp

CMD ["/app/testapp"]

