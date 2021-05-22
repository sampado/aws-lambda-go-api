# aws-lambda-go-api

Basic example of a AWS Lambda & API Gateway with three implementations

## Lambda function handler
https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html
```sh
make run
```
Go to: 
http://localhost:3000/lambda


## Gateway package
https://github.com/apex/gateway package.

```sh
make run
```
Go to: 
http://localhost:8080/gateway?name=sampa

## Gorilla Mux Server

```sh
go run cmd/server/main.go
```
Go to: 
http://localhost:8080/hello?name=sampa


