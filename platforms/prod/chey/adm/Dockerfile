FROM softinnov/cheyenne

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD httpd.adm.cfg /httpd.cfg

ADD esc-adm/src /esc-adm
VOLUME /var/log

EXPOSE 80
