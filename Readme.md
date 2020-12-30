1. First need to install protobuf in your system to access protoc command
    document link http://google.github.io/proto-lens/installing-protoc.html
2. Go packages required for deve;opment
    go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    go get -u google.golang.org/grpc
3. Command for generating go code agains proto file
    protoc --go_out=plugins=grpc:<pkg_name> <proto_file_name>
    Ex: protoc --go_out=plugins=grpc:hello hello.proto
