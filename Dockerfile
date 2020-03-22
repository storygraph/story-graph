FROM scratch
ADD storygraph ./

ENV PORT 8080
EXPOSE 8080
ENTRYPOINT [ "/storygraph" ]