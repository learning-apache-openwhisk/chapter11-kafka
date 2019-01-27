FROM ubuntu:bionic
COPY kat.sh /bin/kat
COPY cred.json /bin
RUN apt-get -y update && \
    apt-get -y install kafkacat jq ca-certificates dos2unix && \
    dos2unix /bin/kat
WORKDIR /bin
ENTRYPOINT ["/bin/kat"]
