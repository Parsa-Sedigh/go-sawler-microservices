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
We make a change to codebase, we need to get that into docker swarm somehow. For this, do:
1. go into the microservice that has some changes and where it's dockerfile lives
2. build a new tagged version of it and push it up:
```shell
docker build -f <name of dockerfile> -t <username>/<name of microservice>:<new version> .
docker push parsa7899/logger-service:1.0.1
```
3. in project folder, update the version of that new image in docker swarm. Now anytime we update a service to a new version, we want to
at least two instances of that service to be running and that's because when we update the service with at least 2 instances running,
there will be almost no downtime. Because it updates the service one at a time for each instance. So we need to scale the service we wanna
update before doing the actual update in swarm:
```shell
docker service scale <service name>=2 # you can get service name from docker service ls
```
4. now update the service to new version:
```shell
docker service update --image <image name with version like parsa7899/logger-service:1.0.1> <service name like myapp_logger-service>

# to verify the new version is up and running:
docker service ls
```
5. update the related `image` of that service to use the new version in swarm.yml

You can downgrade things as well with these steps.

You can automate these things with continuous integration.

## 84-8. Stopping Docker swarm
To bring down a specific service running in swarm:
```shell
docker service scale <service name>=0
```
Doing this for all services running in swarm, doesn't make that swarm to be destroyed. It still exists, it hasn't been removed yet.

If you want to actually remove the entire swarm:
```shell
docker stack rm <app name>
docker service ls
```
```shell
docker swarm leave # for managers, you need to add --force at the end
```

## 85-9. Updating the Broker service, and creating a Dockerfile for the front end
Currently, we have everything in swarm except the frontend and that makes no sense. To put it there, we need to build a version of the 
frontend as a docker image and push it and ... .

Change the port of docker container for broker service to be 8080 instead of 80, so we would have: "8080:8080" and you'll see why when we put
a webserver as our single point of entry into the swarm.

After doing changes in for example the broker service, run these in project folder:
```shell
make build_broker

# build and tag the new docker image. For this go to broker-service folder and run:
docker build -f broker-service.dockerfile -t parsa7899/broker-service:1.0.1 .
docker push parsa7899/broker-service:1.0.1
```
Then in swarm.yml change the version of that image.

We didn't specify a port for broker service, since we won't expose it to internet. Instead, we'll be proxying from caddy or nginx
to the microservice(we could leave that a port 80 but we wanted to make sure things are absolutely clear).

In front-end , we wanna move the `templates` folder into a readonly file system. For this, add special comments(go:embed) to render function of main.go
of front-end. Then change some code in `render` function there.

## 86-10. Solution to the Challenge
Create `build_front_linux` command in Makefile. Then in project folder, run:
```shell
make build_front_linux
```

## 87-11. Adding the Front end to our swarm.yml deployment file
After pushing the frontend docker image to dockerhub, add it to swarm.yml .

Currently, we can bring up the frontend service but we have no means of accessing it(there's no `ports` for front-end and broker-service in
swarm.yml). So the next step is to put a proxy webserver and it will take all reqs for the frontend and broker which are the things that
are gonna be hit and it will route them to the appropriate microservice. 

## 88-12. Adding Caddy to the mix as a Proxy to our front end and the broker
We can't hit frontend and broker because there's no ports exposed(since we don't want to expose them to internet). For this we need to add
another image to our swarm which is a webserver(reverse proxy).

Caddy handles installation and deployment of SSL certificates automatically. As long as we have a registered domain name and appropriate
entry in the name servers, it uses let's encrypt to handle SSL.

Add `caddy.dockerfile` and `Caddyfile` to `project`.

With `Strict-Transport-Security max-age=31536000;`, we don't let people visit the encrypted version of our site. If they try to, just
redirect them to encrypted one.

Now we need to build and tag a docker image for caddy and push it to dockerhub(in project folder where caddy.dockerfile exists):
```shell
docker build -f caddy.dockerfile -t parsa7899/micro-caddy:1.0.0 .
docker push parsa7899/micro-caddy:1.0.0
```

Now add the caddy service in swarm.yml . We need to specify ports for it, because that's how we're gonna access our swarm.
We also need to specify a volume for caddy container. Because when we deploy it to a live server, we're gonna be using port 443 and this
means that caddy has to have somewhere to store the certificates that it uses. Now we **could** for a little while, if we just didn't specify
any volumes and we just say: Ok, I'll put my certificates I'm getting from let's encrypt, right in the docker image and the problem with this is
docker images in a swarm or k8s, these are ephemeral. They don't live forever. They might disappear and come back and eventually as these
docker images disappear and get to re-initialized and re-deployed, it'll keep requesting certificates over and over and let's encrypt
has request limits. So instead, we'll specify some volumes for caddy.

Now we can hit the frontend.

## 89-13. Modifying our hosts file to add a backend entry and bringing up our swarm
In `Caddyfile`, we're specifying two domain names or virtual hosts, one is localhost:80 (make sure you include the port 80 or otherwise it
will try to fetch an SSL certificate using let's encrypt and that won't work) and another one called backend:80 . Now localhost is not a problem,
it's available on any machine. But we need to specify `backend`. For this, run:
```shell
sudo vi /etc/hosts
```

Anytime you try to go somewhere, your computer will always look in your `hosts` file first to see if there's an entry and if there's not,
then it goes out and questions a name server and says: Give me the IP address for the host I'm looking for.

Modify this line: `127.0.0.1       localhost` to: `127.0.0.1       localhost backend`
and this line: `::1             localhost` to `::1             localhost backend`

Now if you run:
```shell
ping backend
```
it will return you localhost(`backend` is now pointing to `localhost`).

Now we wanna bring up docker swarm. First make sure nothing is running, so in docker dashboard's containers tab see if anything is running.

Then run this in `project` folder:
```shell
docker swarm init
docker stack deploy -c swarm.yml myapp
```

Now in web browser, go to `localhost`. 

We can't make any req currently from frontend, because it's making a req to port 8080, but we're not listening on port 8080 anymore,
instead we should send req to: **`http://backend/handle` which will proxy the req to broker-service(why proxying it? Because broker service
doesn't have any ports exposed, so we need a proxy for that**).

Note: We don't want to send req to localhost:80 because that's the frontend! That's what's displaying the page we're on it in the first place!!!
Instead, we wanna go to `http://backend/handle` which we set it up on `hosts` file.

We can hard code the new url, it's gonna work for a short time, as soon as we deploy this microservice to an actual production server on internet,
that's gonna have a different URL.

For this, we need to pass data(from environment) to that go template and that will make it more portable.

### 13.1 Modifying hosts on Windows 1011

## 90-14. Challenge correcting the URL to the broker service in the front end

## 91-15. Solution to challenge
In project folder:
```shell
# first build the binary of frontend:
make build_front_linux

# then in front-end service, build and tag a new docker image of that service:
docker build -f front-end.dockerfile -t parsa7899/front-end:1.0.1 .
docker push parsa7899/front-end:1.0.1
```
Then change the image version in `swarm.yml`. Then add `BROKER_URL` environment variable in swarm file.

Then bring up docker swarm:
```shell
docker stack rm <app name>

# in the folder where swarm.yml exists:
docker stack deploy -c swarm.yml myapp
```

Now test things on `localhost`.

## 92-16. Updating Postgres to 14.2 - why monitoring is important!
If you use postgres:14.0 and deploy it to prod like linode, 100% of CPU would be taken by the node that has postgres installed on it!
The reason for this is kdevtmpfsi process and even if you kill it, it starts back up. Turns out there's some crypto current mining software malware.
So we would helping someone mine cryptocurrency somewhere on the planet! That's not good. To fix this, delete that virtual private machine,
re-install it and update the postgres at least on 14.2 version.

## 93-17. Spinning up two new servers on Linode
The first step is to determine what kind of servers do you wanna use? Because with docker swarm, you can use a virtual private server,
you can use a dedicated server.

We wanna create two new servers. So click on `Create Linode` button. In `Choose a distrubtion` image, select `ubuntu 20.04 LTS`.
For region, you can choose whatever region you want but make sure you put both of servers in the same data center. Although you can create
docker swarms that live in different data centers, it's not recommended because of network latency(put it in region where it's close to you).

Instead of `Dedicated CPU`, use `shared CPU` and choose $10 that has 2GB memory(RAM) and 50GB storage.

For Linode label, for first one, write `node-1` and you can choose the SSH keys checkbox, then `create Linode`.

Then go to `Linodes` page(digital ocean calls them `droplets`). Then create another node and name it `node-2`.

Now we need to configure the servers that we created.

### 17.1 DigitalOcean
### 17.2 Linode
### 17.3 Vultr

## 94-18. Setting up a non-root account and putting a firewall in place.
Non-root means it doesn't have permissions without using `sudo`.

SSH into the server:
```shell
ssh root@<ip address>
```

First thing to do is to add a user:
```shell
adduser <name of user>
# Then give it a password
```

Now we've created this user, we want to give that user sudo privileges so that user can execute commands as root:
```shell
usermod -aG sudo <username>
```

Now we can SSH into the server as that user and we can execute commands by typing `sudo`.

Next thing is to setup a basic firewall and we do this in ubuntu by using `ufw`(uncomplicated firewall):
```shell
ufw allow ssh
ufw allow http
ufw allow https
```
So now we can SSH into and pass the firewall(once we turn it on) and when the webserver is running, we can connect using port 80(http) and
port 443(https). 

Now we need specific firewall ports to be open for docker swarm:
```shell
ufw allow 2377/tcp
ufw allow 7946/tcp
ufw allow 7946/udp
ufw allow 4789/udp
ufw allow 8025/tcp # for hitting mailhog
```

Now let's enable the firewall:
```shell
ufw enable

ufw status # you see firewall is setup for ipv4 and ipv6
```

Now if you ssh into server, with the new created user and run a command with `sudo`, everything should be fine. 

## 95-19. Installing Docker on the servers
The approach we're gonna take is to install docker using the repository. For this, SSH into server and run:
```shell
sudo apt update # it will say some of the packages can be updated, for this, run: sudo apt upgrade

sudo apt-get install \
ca-certificates \
curl \
gnupg \
lsb-release

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

Now install the docker engine:
```shell
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

which docker # prints /usr/bin/docker . So the docker command is now available to us
```

We need to repeat these steps for each of the servers.

### 19.1 Install Docker Engine on Ubuntu

## 96-20. Setting the hostname for our server
We want to set the hostname for the machine we're SSHed into and we want to name it name it `node-1`:
```shell
sudo hostnamectl set-hostname node-1
```

Now if you disconnect by running `exit` and reconnect again using SSH, you notice that before, the prompt was: `<username>@localhost` but
after changing the hostname and logging in again, the prompt is now: `<username>@<hostname>`.

Setting the hostname is important so that docker can find the appropriate servers.

Now let's put a couple of entries in the `hosts` file:
```shell
sudo vi /etc/hosts
```
Put the ip address of current machine in hosts file, so sth like: `<ip address> (hit tab here) <hostname of this machine like node-1>`.
Then put the ip address and hostname of the other servers you have(which should be in the same datacenter): 
`<ip address of your another server> (tab here) <hostname of other server like node-2>`.

Doing this, makes things a lot faster because it doesn't have to do DNS lookups and if the nameserver disappears for whatever reason,
things will still work(although there's no nameserver for getting the IP of a host). It's always a good practice to do this.

**Now do the same thing on the other servers.**

In the next vid, we'll put some DNS entries(nameserver entries) in for our servers and this presupposes that you own a domain name.

## 97-21. Adding DNS entries for our servers
Add DNS entries to point to our servers.

On godaddy.com get a domain. Then in `products` page click on DNS button(you can run your own nameservers by the way). Then on DNS management page,
click on `Add` button. Now add a `A` record and for name, write `node-1` and value should be a valid IP address which you can get this value
from the related server on linode. Leave TTL as default. Then add another `A` record for your other server and name it `node-2`.

Now add another `A` record(you could use a CName record) and name it `swarm`(so we can hit our docker swarm by going to swarm.<domain>) and for value,
use the IP address of your first server. 

Now we add another `A` record for swarm named `swarm` again, but this time, we use the IP address of our another linode server.

Why did we add two `A` records with two separate IP addresses?

When we deploy our docker swarm, we'll be able to go to any node in that swarm using the address `swarm.<domain on your godaddy>` and even if the
webserver is running on a different node, it will still bring the application up and this is one of the nice features about docker swarm, which
is we can hit any node in swarm and it's just like hitting a node that has the service I want. In our case, we'll be running only one instance
of Caddy(webserver) and even if it's running on node-1(which is the case actually), even if my req goes to node-2, I'm still going to have
access to Caddy.

Note: Some nameservers can be slow to update. You can verify it by(make sure you're not logged in to any server) and run:
```shell
ping swarm.<your domain>
ping node-1.<your domain> # I guess
ping node-2.<your domain>
```

Now we wanna get our swarm up and running on both of our servers.

## 98-22. Adding a DNS entry for the Broker service
![](img/98-22-1.png)

Since our frontend app makes req to backend(our broker service), we need to add a DNS entry.

Add CName record and for it's name use `broker`(so frontend would send req to broker.<domain>) and it has to go to swarm.<domain>(as value field).

## 99-23. Initializing a manager, and adding a worker
Open 2 terminal windows, one for each server and on each, connect to the related server(SSH) using non-root user accounts.

To initialize our docker swarm:
```shell
# since we're not working on a single node swarm, we need to add 
sudo docker swarm init --advertise-addr <IP address of the server you're connected to like node-1>
```
The above command will give you a command like: docker swarm join --token ... . Copy it and on the window that is connected to your other server,
type: `sudo <past the command>`.

With this, we have initialized a docker swarm with one manager on node-1(we used IP address of node-1 for advertise-addr option) and 
one worker node in node-2.

Currently, nothing is running(`sudo docker service ls`), because we haven't deployed anything yet.

## 100-24. Updating our swarm.yml and Caddy dockerfile for production
The caddy.dockerfile creates an image but we have a problem. In Caddyfile it's referring to localhost:80 and backend:80 and that's not gonna
work. To fix this, create a Caddyfile.production and modify it to make it work on a cluster. Change localhost:80 to <domain> which would be like:
swarm.<continue> and we leave the :80 there for now, because if we remove that part, when we deploy this, it would try to fetch an SSL certificate
and we wanna make sure everything works before we actually bother installing SSL. 

Then change `backend:80` to <CName of broker like broker-<rest of the domain>> and leave the port.

Then create `caddy.production.dockerfile`.

Then in project folder:
```shell
docker build -f caddy.production.dockerfile -t parsa7899/micro-caddy-production:1.0.0 .
docker push parsa7899/micro-caddy-production:1.0.0
```

Now just as we created a copy of dockerfile for caddy image and the Caddyfile for production, we're gonna do the same thing with swarm.yml .
So create `swarm.production.yml` and there, use the production dockerfile of caddy for it's service.

Now since we have a `volumes` section in our swarm definition, we should remember when we deploy the stack(`docker stack deploy ...`) on the
production servers, we'll have to do it in the same folder where the folders specified in `volumes` section exist which means we need to create
them.

So SSH into node-1 and in root level of that server(`/`). You can put it whereever you want:
```shell
sudo mkdir swarm

sudo chown <username>:<username> swarm/
cd swarm
mkdir caddy_data
mkdir caddy_config
vi swarm.yml # copy swarm.production.yml there
```

Now let's deploy the swarm(you have to be in the folder that has those caddy folders and the swarm.yml):
```shell
sudo docker stack deploy -c swarm.yml myapp
```
Now give it some time to bring everything up. Because it takes a while to initialize the DB firs time around.

## 101-25. Trying things out, and correcting some mistakes
Go to sth like: `swarm.<domain>`. If it's ok, it means the swarm is running and the caddy and front-end microservices are OK.

Add you user to docker group, just so we don't have to keep typing sudo everytime we execute a command for `docker`:
```shell
sudo usermod -aG docker <user>
```

```shell
docker node ps
```

Now in swarm directory, create the folders for docker volumes that we missed for postgres and mongo:
```shell
# in swarm directory
mkdir db-data
mkdir db-data/mongo
mkdir db-data/postgres

# Now stop the swarm since it's not running properly
docker stack rm <name>
```

Another problem is some of our services that are defined as `global`, are running with multiple instances, but on our local dev env, they were
ok which means keep one instance of this service running, on every node of the swarm and that's not what we want. We want `mode: replicated` with
`replicas: 1` for rabbit, mailhog and mongo, so change this in swarm.production.yml .

Another problem is in mongo service, we have a volume for mongo which means we tie a local folder on server to a folder in container and
that'll work fine unless mongo gets moved at some point to a different node, one that doesn't have that directory. To fix this,
add `placement` with constraints for mongo and also for postgres and caddy.

Now put the new content of swarm.production onto the respective file on server. Then deploy the swarm again.

To get updates in real time:
```shell
watch docker node ps

docker service ls
```

## 102-26. Populating the remote database using an SSH tunnel
Let's connect to postgres running in docker swarm. But tutor never does this. He almost never run a DB in a swarm or even in k8s. Instead
he connects to some remote service.

In DB client, connect to a remote server(use the IP of server as host) and turn on the `SSH tunnel` because we're exposing port 5432 from our
docker swarm, we're just blocking it using the firewall.

Create a DB called users and then run the `users.sql` to init the tables and data.

## 103-27. Enabling SSL certificates on the Caddy microservice
We're already listening on port 443(https) on caddy and we already have defined the mounted volumes(caddy_data and caddy_config) to store
the SSL certificates.

Update front-end service BROKER_URL env var to have https instead of http.

Delete `:80` from first virtual host(swarm.<domain>) and second one which is gonna tell caddy: hey, we're gonna be using
auto-generated and auto-managed SSL certificates also `import security` in `swarm.<domain>`.

Now we need to build and tag a new version of caddy docker image and then update the docker image version of caddy in swarm.production :
```shell
# In project folder
docker build -f caddy.production.dockerfile -t parsa7899/micro-caddy-production:1.0.1 .
docker push parsa7899/micro-caddy-production:1.0.1
```

Update the swarm.yml in manager node with swarm.production .

Now:
```shell
docker pull parsa7899/micro-caddy-production:1.0.1
```

If we were updating just one service, we **could** run:
```shell
docker service scale myapp_caddy=2
```

After it scaled it, we update the image:
```shell
# myapp_caddy is the service name
docker service update --image parsa7899/micro-caddy-production:1.0.1 myapp_caddy
```
The above command will update all of the services, one at a time, so we don't get downtime.

This is great for updating **individual** service. But since we changed swarm deployment file on our server(in this case we changed
an env var in front-end service), we need to bring down the swarm:
```shell
docker stack rm myapp # after this, give it a couple of seconds to ensure everything goes down
```

Now we deploy it again and this will use the new env var. This is a thing you're not going to do in the middle of the day, you'll do it
late at night when nobody's on the server and you might give people warning beforehand:
```shell
docker stack deploy -c swarm.yml myapp
```

Now let's see what service is running on what server(we have multiple servers right?). The below command, shows us on the current node,
which services are running:
```shell
docker node ps
```

To see services running on a specific node:
```shell
docker node ps node-2
```

Go to swarm.<domain> to see https and test things.

Currently, there's a problem. Currently, we have a constraint on caddy, postgres and mongo that ensures those services are always going to be
deployed(even if we have multiple instances of each one of them) on a specific node and the reason for doing this, is volumes for those services.
It would be nice if we would have that information(the data that is in the folders specified for volumes), available on every node and if we're
able to do that, then we'd be able to deploy those services on any node we want.

There are a couple of ways to do that. One is `GlusterFS`. If we install it on our master node and all our worker nodes and share a particular volume
or disk directory from one node to others, anytime a file goes into the directory on one node, it gets automatically copied to all nodes.

An alternative approach would be to use `SSHFS` which is an SSH file system, you can mount remote filesystems over SSH. It's a bit slower
than Gluster but in our case, everythig is in the same datacenter and we're not talking about copying huge amounts of data. So it's not
that much different.

### 27.1 GlusterFS
### 27.2 sshfs