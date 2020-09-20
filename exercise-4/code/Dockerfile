FROM debian:stretch

ADD karma /opt/
ADD ui/build/static /opt/static/
ADD ui/build/index.html /opt/static/
WORKDIR /opt

USER 1000
ENTRYPOINT ["/opt/karma"]
