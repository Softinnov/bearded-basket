FROM softinnov/minicheyenne

ADD httpd.cfg /httpd.cfg
ADD mysql-protocol.r /libs/mysql-protocol.r

ADD esc-pdv/src /esc-pdv
VOLUME /var/log

ADD mksh-static /bin/sh

EXPOSE 80
