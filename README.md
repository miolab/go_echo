# Go Echo Sandbox

## About

This is a simple web server application built with the __[Echo](https://echo.labstack.com/)__ framework in __Go__.

## Usage

### Running Locally

Ensure you have [Go](https://go.dev/dl/) (version 1.20 or later) installed.

```sh
go run main.go
```

### Running Locally with Docker

Build the Docker image:

```sh
docker build -t go-echo .
```

Run the Docker container:

```sh
docker run --rm -p 1323:1323 go-echo
```

### Result Confirmation

Using a tool like `curl` to access http://localhost:1323 .

```sh
curl http://localhost:1323
```
