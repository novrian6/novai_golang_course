1.Install the Go Protocol Buffers Plugin:
 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


2.Install the gRPC Go Plugin:
 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


3.Generate Code: Navigate to the directory where your hello.proto file is located, then run the following command to generate the Go code for gRPC:

 
protoc --go_out=hello --go_opt=paths=source_relative --go-grpc_out=hello --go-grpc_opt=paths=source_relative hello.proto
Note :
- the project path under src is /chapter16-grpc-helloword
- go mod init is chapter16-grpc-helloword
- make sure the pb "chapter16-grpc-helloword/hello"
- and in the project path there is folder hello contains the .pb.go

