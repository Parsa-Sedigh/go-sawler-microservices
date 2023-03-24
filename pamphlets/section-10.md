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


## 110-7. Creating a deployment file for RabbitMQ
## 111-8. Creating a deployment file for the Broker service
## 112-9. When things go wrong...
## 113-10. Creating a deployment file for MailHog
## 114-11. Creating a deployment file for the Mail microservice
## 115-12. Creating a deployment file for the Logger service
## 116-13. Creating a deployment file for the Listener service
## 117-14. Running Postgres on the host machine, so we can connect to it from k8s
## 118-15. Creating a deployment file for the Authentication service
## 119-16. Trying things out by adding a LoadBalancer service
## 120-17. Creating a deployment file for the Front End microservice
## 121-18. Adding an nginx Ingress to our cluster
## 122-19. Trying out our Ingress
### 19.1 Edit Windows hosts file
## 123-20. Scaling services
## 124-21. Updating services
## 125-22. Deploying to cloud services
### 22.1 How to configure SSLTLS on Ingress with k8s
### 22.2 ingess_ssl