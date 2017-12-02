FROM golang:1.9.2
ADD . /code
WORKDIR /code
EXPOSE 8080
RUN go build -o server .
CMD ["go", "run", "/code/server.go"]

