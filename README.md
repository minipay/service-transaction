# Service Transaction

This app is a part of minipay project

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

**!! WARNING !!** This project is using Go [Modules](https://blog.golang.org/using-go-modules) which is the minimum version of installed golang is 1.11. 

### Installing

A step by step series of examples that tell you how to get a development env running
1. Clone this project `git clone https://github.com/minipay/service-transaction.git` outside $GOPATH (for version >= 1.11) or inside $GOPATH/src (for version < 1.11)
2. Run `cd service-transaction`. Copy and rename `sample.config.json` to `config.json` 
3. **(for version under 1.11)** After cloning it into your GOPATH, you need to run this to install all dependencies :
`go mod tidy`
4. Setting up your host, port, and database configuration in `config.json`
5. If the `debug` is true, it will print all log including your error
6. Make sure [migration-tools](https://github.com/golang-migrate/migrate) is installed
7. Running without docker : make sure mysql already installed in your machine. Adjust your `config`. Run `go run main.go`
8. Running With docker : </br>
- Dont change anything in `config.json`. Leave it default </br>
- Run `docker build . -t service-transaction:1.0` to build the image </br>
- Run `docker-compose up` and wait. If no error appear, it means your app is ready </br>

## Built With

* [Negroni](https://github.com/urfave/negroni)
* [Viper](https://github.com/spf13/viper)
* [MysqlDriver](https://github.com/go-sql-driver/mysql)
* [http-router](https://github.com/julienschmidt/httprouter)