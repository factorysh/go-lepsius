FROM debian

RUN apt-get update && apt-get install -y haproxy

CMD /usr/sbin/haproxy-systemd-wrapper -f /etc/haproxy/haproxy.cfg
