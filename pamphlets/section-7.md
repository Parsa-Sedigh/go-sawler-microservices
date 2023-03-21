# Section 7. Communicating between services using Remote Procedure Calls (RPC)

## 63-1. What we'll cover in this section
So far, we've been POSTing JSON between microservices.

When using RPC, if one of the ends of communication is not written in go, this won't work.

We're gonna implement this for communications between broker and logger services.

![](img/63-1-1.png)

## 64-2. Setting up an RPC server in the Logger microservice
Create rpc.go in logger-service.

Anytime you wanna implement an RPC server, you do need a specific type for that. We called it `RPCServer` in logger-service.

After defining the rpc server type, payload type, now, we write methods that we want to expose via RPC with the receiver of our rpc server.

## 65-3. Listening for RPC calls in the Logger microservice
We need to start the RPC server.

## 66-4. Calling the Logger from the Broker using RPC
In broker service, instead of `logEventViaRabbit`, we're gonna create a new func named `logItemViaRPC`. 

As the payload for sending to rpc server, you need to create a type that exactly matches the one that the remote RPC server expects to get.
We named that type `RPCPayload`(exact name that we expect in rpc server).

Note: Any method I want to expose in our RPC server, must be exported(has to start with a capital letter), otherwise it's just not going to work.

## 67-5. Trying things out
To run things, in project folder:
```shell
make up_build
make start
```

Note: If there is an unexpected err of port already in use, you should probably make the already up services using that port, down and then
run that service again(it's already running, so first bring it down ).

In order to use RPC, we have to have go on both the client and the server. This was an alternative method of communicating between microservices.
But with GRPC, we can have services written in **whatever language** you want, so you're not limited to working with just go.