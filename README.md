# API Gateway

## Generate Go Code from Proto
```shell script
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative company.proto
```
