# Section 5. Building a Mail Service

## 40-1. What we'll cover in this section
The mail service shouldn't communicate with internet at all. Now it in our test app, it will, since we're gonna create a button
that will communicate directly to the broker and then to the mail service, but typically you're not gonna do that.

You want your mail service to be protected. So it never actually communicates directly with the clear internet.
In other words, you keep it right in your docker swarm or k8s cluster or ... .

But we wanna verify that it works.

Also we want a mail server. In this course, we're gonna communicate with a service called mailhog which is an app that we'll add to
our docker compose, instead you can just use your gmail account or ... . But using gmail or ... is not good, instead use mailhog. Because
it's a good practice when you're in development, not to actually send email to a real mail server like gmail.

## 41-2. Adding Mailhog to our docker-compose
Mailhog simulates a mail server(it won't let emails go to the clear internet).

After adding it to docker-compose, run:
```shell
make down
make up # to pull and start mailhog
```

To verify it's running, in browser go to: `localhost:8025`.

## 42-3. Setting up a stub Mail microservice

## 43-4. Building the logic to send email
## 44-5. Building the routes, handlers, and email templates
## 45-6. Challenge Adding the Mail service to docker-compose.yml and the Makefile
## 46-7. Solution to challenge
## 47-8. Modifying the Broker service to handle mail
## 48-9. Updating the front end to send mail
## 49-10. A note about mail and security