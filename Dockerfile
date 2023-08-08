FROM golang:alpine AS build 

ENV GO111MODULE=on
ENV PORT=3000


WORKDIR /app
COPY go.mod .
COPY go.sum .
# COPY .env ./tmp.env
RUN go mod download
COPY . .

RUN go build -o /img

EXPOSE 3000

CMD ["/img"]

FROM alpine:latest 
WORKDIR /
COPY --from=build /img /img
# COPY --from=build .env .env
EXPOSE 3000
CMD ["/img"]



# FROM golang:alpine as build

# RUN apk add --no-cache git

# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64


# WORKDIR /build

# COPY go.mod .
# COPY go.sum .
# COPY .env .
# RUN go mod download

# COPY . .

# RUN go build -o /img

# CMD ["img"]

# FROM alpine:latest
# WORKDIR /dist
# # RUN cp /build .


# COPY --from=build /img /img
# EXPOSE 3000
# CMD [ "/img" ]