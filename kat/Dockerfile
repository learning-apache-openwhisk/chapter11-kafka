FROM ubuntu:bionic
RUN apt-get -y update && apt-get -y install kafkacat jq ca-certificates dos2unix
COPY kat.sh /bin/kat
COPY cred.json /bin
WORKDIR /bin
ENTRYPOINT ["/bin/kat"]
