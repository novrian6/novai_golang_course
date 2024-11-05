Build the proto file to generate go file
protoc --go_out=plugins=grpc:. book.proto           
