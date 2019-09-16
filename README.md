# Service Transaction

This app is a part of minipay project

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

**!! WARNING !!** This project is using Go [Modules](https://blog.golang.org/using-go-modules) which is the minimum version of installed golang is 1.11. 

### Installing

A step by step series of examples that tell you how to get a development env running
1. Clone this project `git clone https://github.com/minipay/service-transaction.git` outside $GOPATH (for version >= 1.11) or inside $GOPATH/src (for version < 1.11)
2. Copy and rename `sample.config.json` to `config.json` 
3. **(for version under 1.11)** After cloning it into your GOPATH, you need to run this to install all dependencies :
`go mod tidy`
4. Setting up your host, port, and database configuration in `config.json`
5. If the `debug` is true, it will print all log including your error
6. Run the database migration in your terminal/cmd using [migrate](https://github.com/golang-migrate/migrate)
7. Run `go run main.go` or just `docker-compose up` with default config

## Built With

* [Negroni](https://github.com/urfave/negroni)
* [Viper](https://github.com/spf13/viper)
* [MysqlDriver](https://github.com/go-sql-driver/mysql)
* [http-router](https://github.com/julienschmidt/httprouter)