FROM softinnov/cheyenne

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD httpd.caisse.cfg /httpd.cfg

ADD esc-caisse/src /esc-caisse
VOLUME /var/log

EXPOSE 80
