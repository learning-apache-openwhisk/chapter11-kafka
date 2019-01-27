#!/bin/bash
BROKERS="$(jq -r '.kafka_brokers_sasl|join(",")' <cred.json)"
USER="$(jq -r .user <cred.json)"
PASS="$(jq -r .password <cred.json)"
kafkacat \
  -b "$BROKERS" \
  -X sasl.username="$USER" \
  -X sasl.password="$PASS" \
  -X sasl.mechanisms=PLAIN \
  -X ssl.ca.location=/etc/ssl/certs \
  -X security.protocol=sasl_ssl \
  $*
