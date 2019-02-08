# Install protoc compiler
download from https://github.com/protocolbuffers/protobuf/releases

# Install compiler for go
go get -u github.com/golang/protobuf/protoc-gen-go

# Install compiler for python
pip install grpcio-tools

# To generate python client, go to `protos` folder and run this command:
python3.6 -m grpc_tools.protoc -I. --python_out=../pkg --grpc_python_out=../pkg tech_share.proto

# To generate go client, go to `protos` folder and run this command:
protoc -I.  --go_out=plugins=grpc:../pkg tech_share.proto