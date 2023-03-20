# Section 6. Building a Listener service AMQP with RabbitMQ
## 50-1. What we'll cover in this section
Right now, the way that all of our microservices communicate with each other is either:

having the broker receive a req, send it off to a microservice, get a res and send it back

Or

we can communicate directly with for example the auth microservice.

We're gonna implement a listener service and that listener service talks to rabbitMQ which is a server that manages queues.

Now for example a user makes a req to the broker to say: "Hey, I want to authenticate" and all the broker does, it doesn't communicate directly with
the auth service, instead it just pushed instructions into rabbitMQ(AMQP server) and rabbitMQ takes that req and add it to the queue and the
listener pulls to that queue or looks to that queue and says: are there any instructions? and it pulls on out, decides what to do with it based
upon the content that it finds in the queue and that calls the appropriate service.

![](img/50-1-1.png)

## 51-2. Creating a stub Listener service
Create `listener-service` folder.

We need a driver to communicate with rabbitMQ. So install: `github.com/rabbitmq/amqp091-go`

The listener service is not listening as an API or a web service. So we won't have a `cmd` folder, instead we just create a main.go right at the
top level of that service.

The listener service is not gonna periodically connect to the queue and listen for things this way, instead the queue will push it right to us.
So we'll listen to certain queues and we'll specify those in the connection to rabbitmq and anytime there's an event there, we get it directly
from the queue.

Before connecting to rabbit, we need to add it to our docker-compose and actually have it start up and then we can try connect to it.

## 52-3. Adding RabbitMQ to our docker-compose
5672 is the default port for rabbitMQ.

After adding rabbitmq to compose, run:
```shell
docker-compose up -d
```

### 3.1 RabbitMQ on Docker Hub

## 53-4. Connecting to RabbitMQ
After starting the rabbitMQ container, it will take a bit before it's actually initialized.

Note: Often, when we run `docker-compose down`, rabbitMQ refuses to quit. We can get around that by stop the docker itself by quiting
the docker dashboard.

RabbitMQ is slow to start up. So we're gonna write a backoff routine that will attempt to connect, a fixed number of times and what we're gonna
do is slightly different than we did when we connected to postgrs, because it can take a while.

To test things out, in listener service:
```shell
go run main.go
```

## 54-5. Writing functions to interact with RabbitMQ
Broker pushes sth onto the queue. But the listener service(at least right now), is not gonna put anything onto the queue, instead the queue will
push to this service which listens for things.

Create a new package named `event` because we're dealing with events from the queue.

## 55-6. Adding a logEvent function to our Listener microservice

## 56-7. Updating main.go to start the Listener function
After this lesson, let's modify the broker service to push sth into the queue, so the listener service can get it and act upon it.

## 57-8. Change the RabbitMQ server URL to the Docker address
Instead of `amqp://guest:guest@localhost` as connection url, use: `amqp://guest:guest@<service name of rabbitmq in docker-compose>`.

## 58-9. Creating a Docker image and updating the Makefile
Run:
```shell
make up_build
```

## 59-10. Updating the broker to interact with RabbitMQ


## 60-11. Writing logic to Emit events to RabbitMQ
## 61-12. Adding a new function in the Broker to log items via RabbitMQ
## 62-13. Trying things out