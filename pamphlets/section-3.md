# Section 3. Building an Authentication Service

## 18-1. What we'll cover in this section
We want the frontend try to authenticate through broker. The broker calls the authentication microservice to determine whether or not
that user is able to authenticate and then sends back the appropriate res.

We don't necessarily have to use the broker service to authenticate. We could contact the auth service directly, we have to have
it's port exposed to the internet. But only using the broker service to authenticate also works(we could expose them but we need
to use a firewall).

Initially we'll use the broker service as the single point of entry.

## 19-2. Setting up a stub Authentication service
Create `authentication-service` folder and add it to workspace and there, `go mod init authentication`

We can listen again on port 80 for auth service even though the broker service is also listening on port 80, because docker lets multiple
containers listen on the same port and treats them as individual servers.

## 20-3. Creating and connecting to Postgres from the Authentication service
We need to add postgres to our docker-compose.yaml and we need to make sure that it's available before we return the database
connection in `openDB` function. Because the authentication service might start out before the database service does.

Now we need to add postgres to our docker-compose to set up an entry for authentication service and set the appropriate env variable named DSN.

### A note about PostgreSQL
An important note about Postgres Version

In the next lecture, I'm going to ask you to add Postgres to your docker-compose.yml, and I suggest that you use something like this:

```yaml
postgres:
    image: 'postgres:14.0'
    ports:
      - "5432:5432"
    deploy:
      mode: replicated
      replicas: 1
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: users
    volumes:
      - ./db-data/postgres/:/var/lib/postgresql/data/
```

Please ensure that you use version 14.2 of the Postgres image, so that your file looks like this:

```yaml
postgres:
    image: 'postgres:14.2'
    ports:
      - "5432:5432"
    deploy:
      mode: replicated
      replicas: 1
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: users
    volumes:
      - ./db-data/postgres/:/var/lib/postgresql/data/
```

There is a problem with version 14.0, which I'll tell you about later on in the course.

## 21-5. Updating our docker-compose.yml for Postgres and the Authentication service
When we set up ports for postgres service in docker-compose, it allows us to connect using our services within docker and
also to connect using our favorite postgres client from our actually computer.

After adding the postgres service in docker-compose, create a folder which will be our volume for postgres container and name it db-data and
inside there create folders that are written in volumes section(before colon) of postgres service in docker. That is where the postgres
docker container will store our postgres db.

Then create the auth service required things in docker-compose and name that service: `authentication-service`.

We specified a port mapping for authentication-service, that way if we want to authenticate directly with the auth service from the client
instead of using broker service, we can do it. We exposed a port that we can connect to.

Now update Makefile to add commands for building the auth service. For this, create a command named `build_auth` and then update
`up_build` to also run `build_auth`.

Now in `project` folder run:
```shell
make up_build

# To stop everything:
make down

# To start:
make up
```

Now in docker dashboard, you should see containers named: postgres:14.2, project_authentication-service and project_broker-service.

Create a new DB called `users`.

## 22-6. Populating the Postgres database
You can install the toolbox package of trevor sawler instead of copying the helpers.go into each microservice.

### 6.1 Beekeeper Studio

## 23-7. Adding a route and handler to accept JSON
We don't give too much info for the reason of error, just say: Invalid credentials. We don't want them to know whether it's the username or
password that is invalid, in case somebody is trying to hack into the system.

### 7.2 tsawlertoolbox
https://github.com/tsawler/toolbox

## 24-8. Update the Broker for a standard JSON format, and conect to our Auth service
Broker will listen for a req from frontend on `/handle`, then fire a req off to the authentication microservice, receive the response from the microservice and
sends some kind of res back to the end user.

Why we named it /handle?

Because it'll handle all reqs. A single point of entry. Doesn't matter what microservice we're trying to deal with, that is the endpoint
that will get called.

Since we're gonna have a single point of entry to the broker service(/handle), that means we're gonna have to have some kind of 
agreed-upon(predictable) JSON format. So create a type for that named `RequestPayload`.

In `RequestPayload`, we're gonna create a different type for each of the possible actions. Like Auth, Mail, Log and ... .

## 25-9. Updating the front end to authenticate thorough the Broker and trying things out
**Note:** After doing some changes, run:
```shell
make up_build
make start # to start the frontend
```