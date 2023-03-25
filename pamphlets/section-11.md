# Section 11. Testing Microservices

## 126-1. Testing Routes
Create setup_test.go in auth service(it must be called like this because it will run before other tests).

`TestMain` should be called exactly that. It sets up the environment.

In api folder of authentication-service, run:
```shell
go test -v .
```

Currently, if we want to write tests for `data` package of authentication-service, we need to have a database running! And that's a no no
when you're writing unit tests. So we need to change the way we access the DB and we'll use the repository pattern.

## 127-2. Getting started with the Repository pattern for our data package
Create repository.go .

We're not gonna use the New func of models.go , so comment it and type Models. Instead create a type named `PostgresRepository`.

Instead of initializing our DB connections using New function, we're gonna return a repository using `PostgresRepository` and the good thing
about it is anything that satisfies the Repository interface, can be substituted by `PostgresRepository` and we're gonna create a function
called `NewPostgresTestRepository`.

## 128-3. Updating our models, handlers, and the main function to use our repository
Replace the received of methods in models.go from `*User` to `*PostgresRepository`.

## 129-4. Setting up a Test repository
We won't need Conn field in PostgresTestRepository, but we put it there for consistency.

In `test_models`, we have to implement every single member of `Repository` interface.

Note: We're not testing the DB when we use this, we're testing for example the handlers.

## 130-5. Correcting a (rather stupid) oversight in models

## 131-6. Testing Handlers
We have our test repository set up and let's write a test for a handler with that.

One way we're gonna test the handler which would not work.

Since we're writing test, instead of response, we use response recorder(`rr`). Because in `logRequest`, we're calling 
a remote service(logger service) which isn't running. So we need to mock it.

## 132-7. Mocking our call to the logger-service for tests
We're gonna modify the code so that's possible to mock that req to logger service.

Create RoundTrip function(the names are important here).

### 7. Mocking our call to the logger-service for tests