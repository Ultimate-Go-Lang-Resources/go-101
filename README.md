go-101
======

My repository to learn Go Language (http://golang.org)

## Install Go (on OSX)

```bash
brew install go
```

## Coding style

```
variableName
```

## Basic Commands

```bash
# download from remote (download into both ./src and ./bin)
go get github.com/hanksudo/go-101/hello

# run program
go run hello/hello.go
go run functions/functions.go

# build program
cd hello
go build hello.go
./hello

# go install to $GOPATH/pkg
go install newmath

# go install to $GOPATH/bin
go install mathapp

# format code syntax
go fmt -w

# documentation
godoc builtin
godoc net/http
godoc fmt Printf
godoc -http=:8080

# others
go version
go env
go list
go list all
go list -json all
```

## Reference

- [Go Documentation](http://golang.org/doc/)
- [Tour of Go](http://tour.golang.org/)
