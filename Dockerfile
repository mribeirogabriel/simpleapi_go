FROM golang:alpine

WORKDIR /usr/src/myapp
COPY . .
RUN go get .
RUN go build -o /usr/src/
CMD ["/usr/src/api"]