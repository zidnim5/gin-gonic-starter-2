<img src="icon.png" align="right" />

# gin-gonic framework starter 2 [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome#readme)

> Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.


## Project Structure

    .
    ├── cmd                      # Source files (alternatively `lib` or `app`)
    ├── internal                 # Internall service (alternatively `config`)
    |     └── postgresdb                
    ├── pkg                      # Package files (include `unit_test`)
    |     └── todo                
    ├── makefile
    └── readme.md


## Getting Started

Simple way to getting started

```ssh
     git clone https://github.com/zidnim5/gin-gonic-starter-2.git
```

```ssh
     go mod init starter
```

```ssh
     go mod tidy
```

```ssh
     make dev
```


## Unit Testing

Simple way to run unit testing

```ssh
     go test ./pkg/todo -v -run TestUnit
```

```ssh
     go test ./pkg/todo -v -run TestMocking
```