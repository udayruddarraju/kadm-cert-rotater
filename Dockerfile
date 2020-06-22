FROM ubuntu:18.04
RUN apt-get update && apt-get install -y ca-certificates openssl curl
RUN curl -fsSL https://get.docker.com -o get-docker.sh
RUN sh get-docker.sh
COPY kadm-cert-rotate /usr/bin/kadm-cert-rotate
