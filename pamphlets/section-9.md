# Section 9. Deploying our Distributed App using Docker Swarm

## 77-1. What we'll cover in this section
One way of deploying our microservices to production servers is docker swarm.

In lot of situations k8s is really overkill, particularly if you don't have a dedicated person or team who works on k8s all day every day.

Docker swarm in contrast, can easily be managed by an individual person and he doesn't have to work on it all the time.

Brett fisher course on docker and k8s.

Docker swarm allows you to deploy one or more instances of a docker image and have docker swarm take care of how those instances are deployed. 

### 1.1 Docker Swarm

## 78-2. Building images for our microservices
Before using docker swarm, we need to build and tag our docker images of microservices and push them somewhere and we need to push them somewhere,
because when we deploy our docker swarm, we need to **pull** those images. The best place for push these for our purposes, is docker hub.

We tag with our username.

In logger-service folder:
```shell
docker build -f logger-service.dockerfile -t parsa7899/logger-service:1.0.0 .
docker push parsa7899/logger-service:1.0.0

# to login docker hub:
docker login # then you can try pushing the image again
```
Do these command for all your microservices images.

## 79-3. Creating a Docker swarm deployment file
Currently, we have all our microservices pushed up as docker images.

In `project` folder, create `swarm.yml`.

We're gonna use exactly the same names for the services in swarm.yml as we did in docker-compose.yml because if we name the services
the same as we did in docker-compose, then we don't need to change the URLS for calling each of those services.

When you specify a port for a service(in docker-compose or swarm or ...), it means it's gonna communicate with outside world.

You don't need to push the rabbitmq or postgres or other external(not our source code) images to your docker-hub because they're already
on docker hub by other developers(they're used by a lot of developers, so we don't need to push them again!).

Rabbitmq requires a volume(OFC someplace outside of container). For development, not specifying a volume is OK, but in production you need a 
fixed volume for prod.

We're gonna be starting postgres and mongo right in the docker swarm as part of the swarm itself in this course. The tutor never do that. Some people do.
He almost always have postgres running as it's own instance on it's own server and he connects from the swarm to postgres and he makes sure
they're in the same data center so latency is not an issue.

We can purchase a managed DB from leenode or amazon or digital ocean.

## 80-4. Initalizing and starting Docker Swarm
In project folder:
```shell
docker swarm init
```

When we're deploying to a cluster of servers, anytime you start up a docker swarm, one node, at least one node, will be the manager and that's what
handles the orchestration of all of our various docker images. When we want to add another node that's not a manager, let's say we're getting lots
of traffic and now we wanna add one more node, we'll spin up that new server, set up it's firewall and we type the command printed out
by running the `docker swarm init` and that printed command, will add a worker to that swarm.

Now if you need to add a second manager(if you have a large cluster, it's good to have multiple managers in case one manager dies),
you'll run:
```shell
docker swarm join-token manager
```

You can always get the command needed for adding a worker and get the token back for this, by running:
```shell
docker swarm join-token [worker | manager]
```
Will give the token.

Now run:
```shell
make stop
make down
```

To deploy docker swarm, in project folder where swarm.yml exists:
```shell
docker stack deploy -c swarm.yml myapp # myapp is the name of our swarm

docker service ls
```

We would expose ports for our mail server and mongo and postgres and other similar services in production, we just make sure firewall
of our server doesn't allow any external traffic to those ports and if we needed to hit mongo or postgres or ... , we would use SSH tunnel.

## 81-5. Starting the front end and hitting our swarm
```shell
make start
```

We want to have a single point of entry and we can have multiple instances of that single point of entry and we can use nginx or caddy to act as
a proxy server.

## 82-6. Scaling services
One of the great things about any container orchestration service is we can have multiple instances of many of our services, but not all of them, for
example the ones that have `global` mode, we can only have one of those. But if we want to scale a service up, we can run:
```shell
docker service ls
docker service scale <name of the service>=<number of instances>
docker service ls # to verify the result
```

Other great thing is if a services dies, docker swarm will create another instance and bring it back up, so it keeps things up and running.
Even if you're running everything on a single server, in other words, you have one node in your swarm, this is a convenient way of having
multiple instances of what you need to be running, running and to ensure they stay up and running.

When you need to update your docker swarm, the process for doing this is easy. 

## 83-7. Updating services

## 84-8. Stopping Docker swarm
## 85-9. Updating the Broker service, and creating a Dockerfile for the front end
## 86-10. Solution to the Challenge
## 87-11. Adding the Front end to our swarm.yml deployment file
## 88-12. Adding Caddy to the mix as a Proxy to our front end and the broker
## 89-13. Modifying our hosts file to add a backend entry and bringing up our swarm
### 13.1 Modifying hosts on Windows 1011

## 90-14. Challenge correcting the URL to the broker service in the front end
## 91-15. Solution to challenge
## 92-16. Updating Postgres to 14.2 - why monitoring is important!
## 93-17. Spinning up two new servers on Linode
### 17.1 DigitalOcean
### 17.2 Linode
### 17.3 Vultr
## 94-18. Setting up a non-root account and putting a firewall in place.
## 95-19. Installing Docker on the servers
### 19.1 Install Docker Engine on Ubuntu
## 96-20. Setting the hostname for our server
## 97-21. Adding DNS entries for our servers
## 98-22. Adding a DNS entry for the Broker service
## 99-23. Initializing a manager, and adding a worker
## 100-24. Updating our swarm.yml and Caddy dockerfile for production
## 101-25. Trying things out, and correcting some mistakes
## 102-26. Populating the remote database using an SSH tunnel
## 103-27. Enabling SSL certificates on the Caddy microservice
### 27.1 GlusterFS
### 27.2 sshfs