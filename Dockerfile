FROM golang:1.14.3-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
ENV PORT 8080

EXPOSE 8080
RUN go build -o main .
CMD ["/app/main"]
