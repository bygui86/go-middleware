# go-middleware
Sample project to explore middleware concept in Golang

## run

### simple

```shell script
# in one shell
cd simple
go run main.go

# in another shell
curl "http://localhost:8080/ishealthy"
```

### sub-router

```shell script
# in one shell
cd sub-router
go run main.go

# in another shell
curl "http://localhost:8080/ishealthy"
curl "http://localhost:8080/sub/a"
curl "http://localhost:8080/sub/b"
```

## links

- https://levelup.gitconnected.com/middlewares-for-golang-web-apps-8742e28eef6e
