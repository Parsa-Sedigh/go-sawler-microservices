# In this dockerfile, first we build all of our code on one docker image and then create a much smaller docker image(because it's just
# an alpine image) and just copy the built executable(named brokerApp) to there.

# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app
COPY . /app

WORKDIR /app

# we're not using any c library, just use the standard library
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# add(+) the executable flag(x) to /app/brokerApp
RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]