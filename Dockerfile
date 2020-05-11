FROM alpine:3.7

RUN apk --no-cache add curl

ADD storygraph ./

ARG CONTAINER_PORT
ENV PORT $CONTAINER_PORT
EXPOSE $CONTAINER_PORT

ENTRYPOINT [ "/storygraph" ]