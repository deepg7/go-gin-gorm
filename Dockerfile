FROM golang:alpine

ENV GO111MODULE=on
ENV PORT=3000

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /img

EXPOSE 3000

CMD ["/img"]