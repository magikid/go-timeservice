FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ADD timeservice /bin/timeservice

CMD ["timeservice"]
