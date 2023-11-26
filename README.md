# connect-go-sample

Ref: [Getting started | Connect](https://connect.build/docs/go/getting-started/)

You can create both http/grpc server using connect-go.

## install tools

```bash
make install-tools
```

## try connect-go http server

```bash
$ go run ./cmd/server

$ curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    -XPOST http://localhost:8080/apis.v1.HealthService/Health
{"status":"OK"}

$ curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/apis.v1.GreetService/Greet

{"greeting":"Hello, Jane"}
```

## add api

```bash
touch ./apis/v1/your_newapi.proto

make all
```
