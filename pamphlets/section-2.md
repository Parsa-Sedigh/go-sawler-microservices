# Section 2. Building a simple front end and one Microservice
## 9-1. What we'll cover in this section
## 10-2. Setting up the front end

## 11-3. Reviewing the front end code
`file>add folder to workspace` and choose front-end folder.

To start the frontend:
```shell
# assuming we're in the frontend folder
go run ./cmd/web
```

Then go to `http://localhost`

## 12-4. Our first service the Broker
Stop the frontend app by hitting: `ctrl + c`

Create a folder at the same level of front-end, named `broker-service`.

Then click on `file>add folder to workspace` and choose broker-service.

Then inside `broker-service`, run: `go mod init`. Then create `cmd` folder there and inside there, another folder named `api`.

The main entry point to the broker-service is inside cmd/api/main .

We want to make sure that we can connect from front-end to the broker-service. broker-service is gonna take reqs and forward them
off to some microservice and then send a res back. For this, install `go get github.com/go-chi/chi/v5` in the broker-service.
Also: `github.com/go-chi/chi/v5/middleware` and `github.com/go-chi/chi/cors`.

## 13-5. Building a docker image for the Broker service
1. multi-stage build using a certified go docker image
2. 

The second way is much faster.

We want a docker-compose file that runs all of our microservices.

Vscode only allows adding folders not files to a workspace, so we can use the finder to create a file next to directories in a workspace.

Create a dir named `project`. After creating it, add it to workspace using: `file > add folder to workspace` in vscode.

Then create `dockerfile`s for each service to tell docker-compose how to build the images of those services.

In docker-compose, in this case, we're only ever allowed to have 1 replica because we can't listen with two or more docker images on 
port 8080 on the localhost. But later on when we implement service discovery, we'll be replicating some of our docker images.
We just want to get used to writing `replicas`, although for now it's only 1 replica.

Now go into `project` folder which is where docker-compose file lives and run:
```shell
docker-compose up -d
```

Now in docker dashboard > containers/apps section, you can see the project running that is named `project`.

Leave this one running and let's continue.

## 14-6. Adding a button and JavaScript to the front end
In frontend folder:
```shell
go run ./cmd/web
```

Then go to localhost:80.

## 15-7. Creating some helper functions to deal with JSON and such

## 16-8. Simplifying things with a Makefile (Mac & Linux)
Put the makefile in the `project` folder and make sure it's called `Makefile`.

To use it, let's make sure docker-compose is not running(run this command in project folder):
```shell
docker-compose down
```
Then run:
```shell
make stop
make up_build
```

Currently, in up_build of Makefile, it builds a brokerApp executable and we're doing the same thing all over again in it's dockerfile too.
So remove the build related commands in it's dockerfile. In other words, we're building the executable twice! Once in the dockerfile and 
once in the Makefile.

Now running the `make up_build` would be faster(make sure you run `make down` first).

If you want to start the frontend:
```shell
make start
# to bring it down:
make down
```

## 17-9. Simplifying things with a Makefile (Windows)
