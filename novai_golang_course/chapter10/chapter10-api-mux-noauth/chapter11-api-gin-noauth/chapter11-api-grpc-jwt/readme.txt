How insta proctoc-gen-go_out
go install google.golang.org/protobuf/cmd/protoc-gen-go


Build the proto file to generate go file
protoc --go_out=plugins=grpc:. book.proto           

