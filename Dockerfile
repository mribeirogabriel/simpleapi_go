FROM golang:alpine

WORKDIR /usr/src/myapp
COPY . .
RUN go get .
RUN go build -o /usr/src/
EXPOSE 8081
CMD ["/usr/src/api"]