FROM golang:1.15-alpine3.12 as build

WORKDIR /go/src
COPY src/ ./

RUN apk --no-cache update \
  && go build -o turtleshell

FROM alpine:3.12
LABEL maintainer="Josh Grancell <josh@joshgrancell.com>"

COPY --from=build /go/src/turtleshell /usr/bin/turtleshell

RUN chmod +x /usr/bin/turtleshell \
  && ln -s /usr/bin/turtleshell /bin/turtleshell \
  && ln -s /usr/bin/turtleshell /usr/bin/tsh \
  && ln -s /usr/bin/turtleshell /bin/tsh \
  && adduser -D -s /usr/bin/turtleshell donatello \
  && passwd donatello -d kowabunga \
  && rm /bin/sh \
  && ln -s /bin/tsh /bin/sh

WORKDIR /home/donatello
USER donatello
COPY docker/.turtlerc ./

CMD 'turtle'
