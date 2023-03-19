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
We could send email just using the standard library, but let's use:
- github.com/vanng822/go-premailer/premailer: This allows us to use CSS and automatically have that css converted to a format that's really good
for email
- github.com/xhit/go-simple-mail/v2


We created the `getEncryption` because it makes our lives easier if we switch mail servers and we will definitely do when we go to production, mailog
is not a prod mail server!

## 43-4. Building the logic to send email

## 44-5. Building the routes, handlers, and email templates
You never want to give a path to send an email, to the internet without some kind of security.

## 45-6. Challenge Adding the Mail service to docker-compose.yml and the Makefile

## 46-7. Solution to challenge
To test things, in project folder:
```shell
make up_build # to build a run the microservices
```
## 47-8. Modifying the Broker service to handle mail
In broker, we need to receive a req from frontend, have the broker process it, send it off to the mail service and then send some kind of 
response back. **Obviously, this is not sth you would do with a live app.** Instead, you would have the microservices talk to each other and
never communicate directly from the frontend to the mail service(even with broker service to send it to mail service). Otherwise, we would
send lots of spam emails.

Before going to prod, change `MarshalIndent` calls to `Marshal`.

## 48-9. Updating the front end to send mail

## 49-10. A note about mail and security