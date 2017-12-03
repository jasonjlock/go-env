# Go Development Environment

A Go development environment using Docker Compose.

## Getting Started

Download and install Docker for [Mac](https://www.docker.com/docker-mac) or [Windows.](https://www.docker.com/docker-windows)

Clone this repository:

```
git clone https://github.com/jasonjlock/go-env.git
```

Change directories to the new repository:

```
cd go-env
```

## How To Use

Start the Docker containers:

> Running this step for the first time builds the container images.
> This process can take a while.

```
docker-compose up
```

this starts Nginx and Go containers. Nginx is a reverse proxy
passing requests to the Go server.

## What Next

While the containers are running, visit http://localhost in your browser.

Create a new package:
```
mkdir your_package
```
add your new go package, compile, and run it:
```
docker exec -it goenv_app_1 go build -o your_package /code/your_package/example.go
```
```
docker exec -it goenv_app_1 go run /code/your_package/example.go
```

Stop the running containers:
```
docker-compose down
```

## Helpful Resources

* [Docker Compose](https://docs.docker.com/compose/reference/)
* [Docker Cleanup Commands](https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes)
* [Nginx Reverse Proxy](https://www.nginx.com/resources/admin-guide/reverse-proxy/)
