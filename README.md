# Card validation api ~~as a test task~~

Using Golang `v1.22.0`

### How to start with Docker

Build the image

```shell
docker build --tag card-validation-api .
```

Run the container

```shell
docker run -p 8080:8080 --name card-validation-api card-validation-api
```

### How to start locally

Run tests

```shell
go test ./... -v
```

Run the server

```shell
go run .
```
