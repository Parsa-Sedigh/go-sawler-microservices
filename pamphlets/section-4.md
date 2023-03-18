# Section 4. Building a Logger Service

## 29-1. What we'll cover in this section
The logger service will have no actual connection to the internet. It's only available to things within our microservices cluster
which might exist as docker swarm or might exist just in docker(like it does right now), or ultimately in a k8s cluster.

## 30-2. Getting started with the Logger service
The logger service doesn't face the internet, it's only talked to by other microservices.

Create `logger-service` folder then add it to workspace.

```shell
go get go.mongodb.org/mongo-driver/mongo

# This package is probably already installed, but we run it anyway(if after running this command, nothing was logged, it means it was
# already installed!): 
go get go.mongodb.org/mongo-driver/mongo/options
```

In `connectToMongo` function, we're setting the authentication for mongo, but it would be better if it's(`options.Credential` fields) specified 
in docker-compose.yaml . So we need to get username and password from environment variables or we'll pass them as command line parameters.

## 31-3. Setting up the Logger data models
We need to disconnect from mongo whenever that service exits for whatever reason. For this, we need a context. So we created a context in the
main function of logger service.

The id in mongo is string type instead of int(like in postgres).

Note: In mongo, if collection does not exist, it'll be created.

In models.go, `Update` doesn't take any parameters because we can get the data we need from the receiver.

## 32-4. Finishing up the Logger data models

## 33-5. Setting up routes, handlers, helpers, and a web server in our logger-service
Now we need to add the mongo service to docker-compose and specify the ports for the local port on our server and in the docker container.

## 34-6. Adding MongoDB to our docker-compose.yml file
We're gonna have 1 instance of mongo, so we don't specify the `deploy` in it's docker-compose definition.

To run mongo:
```shell
make down
make up # will pull mongo and start it up with others(because we just defined mongo service in docker-compose)
```

To test things, change mongoURL in main.go to `mongodb://localhost:27017` instead of `mongodb://mongo:27017` and in the logger-service, run:
```shell
go run ./cmd/api/
```

Now we need to modify Makefile and docker-compose.yml to add the logger-service.
Then some modification in broker-service to allow it to communicate to the logger.

## 35-7. Add the logger-service to docker-compose.yml and the Makefile
We shouldn't specify `ports` section for logger-service in docker-compose. Because we're not gonna exposing a port to our local computer from
that service.

Now run:
```shell
make up_build
```

## 36-8. Adding a route and handler on the Broker to communicate with the logger service
**Ultimately**, every one of our microservices will be able to communicate **directly** with the logger microservice without going through the 
broker.

In production, we wouldn't use MarshalIndent(), we would use Marshal(), but the former makes things readable during development.

## 37-9. Update the front end to post to the logger, via the broker

## 38-10. Add basic logging to the Authentication service

## 39-11. Trying things out
In `project` folder:
```shell
make up_build
make start
```

Then go to localhost and test.

Now we should have entry in mongo. For this, connect to mongo service using a mongo client using sth like **mongodb compass**.
Mongodb compass wants you to have your mongo DB on their cluster.

The connection string is: `mongodb://admin:password@localhost:27017/logs?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false`

Use this string to connect in jetbrains IDE's DB.

Note: You may not need all of the query params in connection string!

Note: To connect to mongo, the related docker containers should be running OFC.

### 11.1 MongoDB Compass