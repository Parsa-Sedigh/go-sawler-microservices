# Section 10. Deploying our Distributed App to Kubernetes 

## 104-1. What we'll cover in this section
We're gonna bring up our k8s cluster using minikube. Another tool we need to have is kubectl.

Now let's deploy a cluster on our local machine.

## 105-2. Installing minikube
It allows you to set up a local k8s cluster and manage it.

He installed both minikube and kubectl using binary link.

### 2.1 Minikube

## 106-3. Installing kubectl

### 3.1 kubectl

## 107-4. Initializing a cluster
In docker dashboard, make sure you have no containers running(see the containers page).

This is not gonna work in first try:
```shell
# should print: Done! Kubectl is now configured to use "minikube" cluster and "default" namespace by default
minikube start --nodes=2

minikube status
```

Now we have 2 containers running(you can see this in docker dashboard too):
- minikube
- minikube-m02

## 108-5. Bringing up the k8s dashboard
Run:
```shell
minikube stop
```
The reason for this is one of the irritating things about minikube is when our system goes to sleep, when we wake it up,
we have to restart everything because minikube just won't come up. So before we finish for the day(before you want your system
goes to sleep), run the above command and when
you come back, run:
```shell
minikube start
```

minikube has k8s dashboard installed by default. To look at minikube(k8s) dashboard, run:
```shell
minikube dashboard
```

## 109-6. Creating a deployment file for Mongo
Make sure minikube is running by running: `minikube status`.

When we were working with docker swarm, the way swarm works is that you deploy one or more containers to the swarm but in k8s, instead of
deploying containers, we deploy pods and a pod can have one or more container or service in it. The basic unit within a docker swarm is
a container(in our case docker containers), but in k8s, the basic unit in a k8s cluster is a pod.
```shell
kubectl get pods
kubectl get pods -A # Also shows us the basic components that make k8s available to us in the first place
```

There are 2 ways to define what a pod consist of:
- imperative way: We just type commands using kubectl and define our various pods. We can use this approach to make changes to an existing cluster
- declarative way: we create some files that describe what our resources are gonna be and then we deploy that file using kubectl. This is
less error prone way

Create k8s folder in `project` and there create mongo.yml .

When we deploy these files, yml is converted by kubectl into json.

**Note:** In production, we will never put mongo, rabbit, postgres and ... in a docker swarm or k8s cluster. Instead, we would 
have a mongo, rabbit or ... service running external to k8s and swarm and connect to that.
For rabbit, we would have an installation of rabbit on another server. 

The ports in template>spec>containers>ports doesn't do anything, it's for other people who have to read that file. So it's purely descriptive.

After defining a deployment, we need to have access to it. For this, we need to define the service that's associated with that deployment we wrote.
We could put that service definition in it's own file.

But a preference is when we define a pod, we can always put it's services in the same file. But it needs 3 dashes.

Now we want to deploy everything that's in the k8s folder. So in `project` folder:
```shell
kubectl apply -f k8s
kubectl get pods
minikube dashboard
kubectl get svc
kubectl get deployments
```

## 110-7. Creating a deployment file for RabbitMQ
It's important to use `rabbitmq` as name in rabbit.yml since we're using `rabbitmq` in reqs that we do in microservices.

In project folder:
```shell
kubectl apply -f k8s/rabbit.yml
```

## 111-8. Creating a deployment file for the Broker service

## 112-9. When things go wrong...
With this command, you can see how many times k8s has restarted services. With this, we can find out we have problems in those services.
```shell
kubectl get pods
```

Copy the pod name you got from previous command for the service that's not running OK and run:
```shell
kubectl logs <pod name>
```

One problem that we have is in broker-service and that is we can't connect to rabbit. Let's see the log files for rabbitmq pod as well... .

There are no logs, so looks like it's OK.

The selector in services of our pods should match the selector defined in pod. So instead of `name: <>`, we should use `app: <>` in selector
of a service.

Now we need to get rid of existing deployments and services and start them again. In project folder:
```shell
kubectl get deployments
kubectl delete deployments <deployment-1 name> <deployment-2 name> ...

kubectl get svc
kubectl delete svc <svc-1 name> <svc-2 name> ...
```

**Note:** Do not delete the `kubernetes` service in svc commands!

We should re-deploy the changed files, so in project folder
```shell
kubectl apply -f k8s
```

By running `kubectl get pods` we see rabbitmq pod is restarted multiple times which means it was crashed multiple times.
It's because we need to increase the resource limits for rabbitmq.

Now re-deploy rabbit.yml (you don't need to stop the deployment and apply it again, you can just apply it when it's running).

## 113-10. Creating a deployment file for MailHog
In mailhog, the 1025 port is for smtp-port(for sending mails) and port 8025 is for it's UI dashboard.

## 114-11. Creating a deployment file for the Mail microservice

## 115-12. Creating a deployment file for the Logger service
**Note:** Make sure the restarts column when running kubectl get pods or in minikube's deployments section stays 0. Because if that number
is going up, sth's not right.

## 116-13. Creating a deployment file for the Listener service
We'll not create a pod for postgres which is a good thing and it simulates the way you might do it in a production environment. We'll
create a postgres service running on our local machine and we'll connect to that remotely. So our cluster will connect to an external
postgres service.

## 117-14. Running Postgres on the host machine, so we can connect to it from k8s
We don't want to put postgres in k8s cluster or swarm. But we wanna simulate the way it would work in a production env. For this,
we would have postgres running as some kind of remote service like a managed DB from digital ocean or ... or we might set up our own
postgres DB on it's own virtual machine on linode.

Don't create postgres.yml in k8s folder which would be a docker-compose file.

In project folder, before running this, make sure the only containers that you're running are related to minikube(`minikube` and `minikube-m02`).
You can verify this in docker dashboard.
```shell
docker-compose -f postgres.yml up -d
```

Now this postgres container should appear in docker dashboard.

## 118-15. Creating a deployment file for the Authentication service
```shell
kubectl apply -f k8s/authentication.yml
```

Then check the logs of authentication service pod:
```shell
kubectl get pods # then get the name of authentication pod
kubectl logs <pod name> # should say: Connected to Postgres!
```

## 119-16. Trying things out by adding a LoadBalancer service
Let's try to hit the broker service in our cluster, by running the frontend locally. But this won't work even though it 
says it's(broker-service) listening on port 8080.

You can see list of ports with:
```shell
kubectl get svc
```

The problem is none of those ports(with the exception of k8s) that are listed in deployment files, are available to the outside world.
For this, we need to expose at least one service from that cluster to the local network and there are 3 different ways of doing that in k8s:
- node port
- load balancer(the one we're gonna do)
- nginx ingress(we use this in production). Much like we used caddy in our swarm file as a means of getting to the services inside our swarm,
we can have sth called a nginx ingress that will do the same thing for our k8s cluster

Let's delete the broker-service and replace it with another service that's a load balancer.
```shell
kubectl delete svc <service name>
kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
```
The port is the port we hit from outside world.

Now if you run `kubectl get svc`, you see the `EXTERNAL-IP` for `broker-service` is `pending`.

Now since we create a load balancer we have to run this command and leave it open:
```shell
minikube tunnel
```
We can't close that terminal window until we're done with that tunnel.

So open a new terminal tab.

In frontend service, change the port we're listening on from 8081 to 80 and change the ports in `BrokerURL` of test.page.gohtml and for 
it's value, use: `http://localhost:8080` to hit the local k8s cluster on port 8080 which is now a load balancer exposing the broker-service.

In frontend:
```shell
go run ./cmd/web
```
Now test things. Means in k8s cluster, we have that tunnel exposing a load balancer to the broker-service, everything is working.

## 120-17. Creating a deployment file for the Front End microservice
Let's change the listening port of front-end service to 8081.

Stop the frontend and tunnel and get rid of it:
```shell
kubectl delete svc broker-service
```

Now re-apply the broker-service.yml:
```shell
kubectl apply -f k8s/broker.yml
```

```shell
kubectl apply -f k8s/front-end.yml
```

Let's setup an ingress for minikube.

## 121-18. Adding an nginx Ingress to our cluster
An ingress is nothing more than a web server called nginx that will handle reqs from outside the cluster, examine them and determine where
they need to go and send them there. It's like Caddy in our docker swarm.

Why don't use caddy in k8s?

There is an experimental ingress based upon caddy. Wait it's at version 1.02 or 1.03 to use it.

Before we can add an ingress to our cluster, we need to enable an add on with minikube. Make sure minikube is running and run:
```shell
minikube addons enable ingress # should print: The 'Ingress' addon is enabled.
```

Now we need to define a deployment file for ingress. Put in `project` folder(not in k8s folder). We put it there,
once everything is ok with this new deployment file. So leave it outside of that k8s folder.



## 122-19. Trying out our Ingress
Make sure minikube and all pods are running and run this in project folder:
```shell
kubectl apply -f ingress.yml
kubectl get ingress
```

Let's edit the hosts file:
```shell
sudo vi /etc/hosts
```

Add this to end of the file: `127.0.0.1       front-end.info broker-service.info`

Now run:
```shell
minikube tunnel
```

Now if you go to localhost, nothing shows up. Instead, go to `http://front-end.info`

Now regardless to which cloud provider you're using for k8s cluster, the steps are same but each cloud provider has it's own quirks.

It's gonna cost to deploy this cluster!

### 19.1 Edit Windows hosts file

## 123-20. Scaling services
To scale things up or down, there are a couple of ways:
- use minikube dashboard and changes replicas
- change the `replicas` in deployment yml file and apply it. It will configure the deployment but the service remains unchanged

Unlike docker swarm, one of the things you can do in k8s, is to configure it to automatically add resources depending upon how much
traffic your site is getting or scale back down if there's not much traffic. So we can have **auto-scaling clusters**.

## 124-21. Updating services
After updating the deployment file and applying it, just like the docker swarm, it updates the pods one at a time which minimizes your downtime.

## 125-22. Deploying to cloud services
You need to set up the necessary virtual servers, then install k8s on servers and do it manually. But nobody does that. We use
kubernetes service of linode for example.

Linode gives you the control plane server for free and you just pay for your actual node pools(some kind of virtual private server).

People deploying to k8s, typically are worried about performance so they go with dedicated CPUs instead of shared one.

You're not gonna deploy your cluster without SSL.

### 22.1 How to configure SSLTLS on Ingress with k8s
### 22.2 ingess_ssl