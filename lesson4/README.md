# GOLESSON/lesson3

I try to create simple grpc server and test this api.


## Initialization 

Optional just first setup: We need to protobuf cli;
https://grpc.io/docs/languages/go/quickstart/

```
brew install protobuf
brew install protoc-gen-go
```

Create a new project;

```
go mod init github.com/misaliperver/golesson/lesson4
go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0
go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
```

create new {servicename}.proto file under proto folder. And lets generate protocol buffer file.

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto
```

## What i need to improve that api?

Because grpc is a new tech for micro services approach. That will be faster than rest api's. Json consume high resourec that network or cpu. encoding is very expensive than grpc. Huge service need to optimization practice and grpc approach answers efficiently.