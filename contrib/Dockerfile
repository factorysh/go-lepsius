FROM  bearstech/golang-dep:latest

USER root
RUN apt-get update && \
        apt-get install -y libsystemd-dev && \
        rm -rf /var/lib/apt/lists/*

USER go
