# Section 4. Building a Logger Service

## 29-1. What we'll cover in this section
The logger service will have no actual connection to the internet. It's only available to things within our microservices cluster
which might exist as docker swarm or might exist just in docker(like it does right now), or ultimately in a k8s cluster.

## 30-2. Getting started with the Logger service


## 31-3. Setting up the Logger data models
## 32-4. Finishing up the Logger data models
## 33-5. Setting up routes, handlers, helpers, and a web server in our logger-service
## 34-6. Adding MongoDB to our docker-compose.yml file
## 35-7. Add the logger-service to docker-compose.yml and the Makefile
## 36-8. Adding a route and handler on the Broker to communicate with the logger service
## 37-9. Update the front end to post to the logger, via the broker
## 38-10. Add basic logging to the Authentication service
## 39-11. Trying things out
### 11.1 MongoDB Compass