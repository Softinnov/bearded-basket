FROM softinnov/minicheyenne

ADD httpd.cfg /cheyenne/httpd.cfg
ADD mysql-protocol.r /libs/mysql-protocol.r

ADD esc-pdv /cheyenne/esc-pdv
VOLUME /var/log

EXPOSE 80
