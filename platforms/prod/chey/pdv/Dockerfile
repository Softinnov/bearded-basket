FROM softinnov/cheyenne

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD httpd.pdv.cfg /httpd.cfg

ADD esc-pdv/src /esc-pdv
VOLUME /var/log

EXPOSE 80
