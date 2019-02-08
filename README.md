# Techshare
Sample Project with REST API and gRPC server


## Installation
### Install protoc compiler
download from https://github.com/protocolbuffers/protobuf/releases

### Install compiler for go
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Install compiler for python
```
pip3 install grpcio-tools
```

## Command
### To generate python client, go to `proto` folder and run this command:
```
python3.6 -m grpc_tools.protoc -I. --python_out=../pkg --grpc_python_out=../pkg tech_share.proto
```

### To generate go client, go to `proto` folder and run this command:
```
protoc -I.  --go_out=plugins=grpc:../pkg tech_share.proto
```

## Running Server
```
go run main.go -mode=<rest or grpc> -port=<optional>
```

###  Sample Run REST server
```
go run main.go -mode=rest
```

### Running gRPC server in specific port
```
go run main.go -mode=grpc -port=11000
```
