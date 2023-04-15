FROM ubuntu:jammy

WORKDIR /app

COPY ./.bin/app .

ENTRYPOINT ["./app"]
