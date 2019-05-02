FROM golang:1.12-alpine3.9 as build

WORKDIR /go/src
COPY src/ ./

RUN apk --no-cache update \
  && go build -o turtleshell

FROM alpine:3.9
LABEL maintainer="Josh Grancell <josh@joshgrancell.com>"

COPY --from=build /go/src/turtleshell /usr/bin/turtleshell
COPY docker/docker-entrypoint.tsh /usr/bin/docker-entrypoint

RUN chmod +x /usr/bin/turtleshell \
  && ln -s /usr/bin/turtleshell /bin/turtleshell \
  && ln -s /usr/bin/turtleshell /usr/bin/tsh \
  && ln -s /usr/bin/turtleshell /bin/tsh \
  && chmod +x /usr/bin/docker-entrypoint \
  && adduser -D -s /usr/bin/turtleshell donatello \
  && passwd donatello -d kowabunga

WORKDIR /home/donatello
USER donatello
COPY docker/.turtlerc ./

CMD ['/usr/bin/turtleshell', 'turtle']
