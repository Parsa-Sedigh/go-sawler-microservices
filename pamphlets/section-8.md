# Section 8. Speeding things up (potentially) with gRPC

## 68-1. What we'll cover in this section
The big difference between gRPC and RPC is that you don't have to have both ends(client and server) written in go in gRPC.

gRPC is not always going to be faster than RPC(minIP experienced a slow down when migrating from RPC to gRPC!), but in many cases, we get
faster, it depends on the use case.

### 1.1 gRPC website

## 69-2. Installing the necessary tools for gRPC
We can share the proto file with another developer that works with another language like Java and he can type the appropriate command in their
Java development environment and it will generate the necessary java code.

Run:
```shell
# To follow the course, install the exact specified version instead of installing the latest version
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
```

1. define a .proto file
2. use the two tools we installed to compile it by running a command. Then we have some code to work with
3. write the server code
4. write the client code
5. test things out

## 70-3. Defining a Protocol for gRPC the .proto file
Create logs folder in logger-service.

## 71-4. Generating the gRPC code from the command line
You can install `protobuf support by peterj` in vscode.

Install `protoc`. 

Tip: If you want to install it from precompiled binaries, do not install `protobuf-all...`, because that's too much stuff.
On mac install the apple silicon one and that would be sth like: protoc-<version>-osx-aarch... .

After downloading it, put the downloaded /bin directory in somewhere that's available in PATH, so we can just type protoc and run it.
We can put it in go/bin directory because that directory is the one that regardless of your OS, where your go binaries are installed:
```shell
# protoc is the name of the binary we wanna copy
cp protoc ~/go/bin/

which protoc # would found it

# But now if we try to execute protoc it's gonna prevent us from that. Go to security & privacy>general>>unlock the lock>allow anyway
protoc --version
```

In the folder that has .proto file, run:
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```

### 4.1 Protocol Buffer Compiler Installation

## 72-5. Getting started with the gRPC server
Create grpc.go in logger-service.

Install these in logger-service:

```shell
go get google.golang.org/grpc
go get google.golang.org/protobuf # might have been installed by the last command, but let's run to make sure
```

The `logs.UnimplementedLogServiceServer` field is to ensure backwards compatibility.

## 73-6. Listening for gRPC connections in the Logger microservice

## 74-7. Writing the client code
In broker-service, create `logs` folder and copy yhe logs.proto in logger-service and paste it there and go to logs folder and run the
protoc command mentioned in listen 4 to generate the same source code that exists on the server.

Then install grpc package we mentioned in lesson 5 in broker-service.

## 75-8. Updating the front end code


## 76-9. Trying things out
In project folder:
```shell
make up_build
make stop
make start

# wait for rabbitmq service to start up. You can go to docker-dashboard and then to project-broker-service-1 to see when it starts up
```

gRPC is great for communicating between microservices but currently there are 2 experimental web browsers that communicate between the browser and
some remote service using gRPC. So currently, we can't really use it in a web browser.