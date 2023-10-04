# gRPC 
----
## Protocol Buffer commands

### Manual Code Generation
```protoc -Igreet/proto --go_out=. --go_opt=module=github.com/MartinLupa/gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/MartinLupa/gRPC greet/proto/dummy.proto```

Explanation of command flags:

- ```-Igreet/proto```: -I/<path> indicates where the input file is.

- ```--go_out=.``` and ```--go-grpc_out=.```: indicate that we want to generate the code in the current directory. But will be influenced by ```--go_opt``` and ```--go-grpc_opt```.

- ```--go_opt=module=github.com/MartinLupa/gRPC``` and ```--go-grpc_opt=module=github.com/MartinLupa/gRPC```: we are indicating that our module is called "github.com/MartinLupa/gRPC", so the command will remove this part of the string from the go_package(originally "github.com/MartinLupa/gRPC/greet/proto") and will leave only "./greet/proto", so the output of the command will be available in such path of ./greet/proto.

- ```greet/proto/dummy.proto```: indicates where the .proto file is located.

### Code Generation with Makefile

- ```make <project_name>```: to generate the code out of the .proto file.

- ```make help```: to see all available options.



## Projects

Check the README.md file inside each project to learn how to use it.

### greet: gRPC Unary

Simple greeting gRPC Unary API.

[greet: gRPC Unary README.md](./greet/README.md)

