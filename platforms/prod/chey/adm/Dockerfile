FROM softinnov/cheyenne

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD httpd.adm.cfg /httpd.cfg

RUN mv /cheyenne /chey_adm

ADD esc-adm/src /esc-adm
VOLUME /var/log

EXPOSE 80

CMD ["/chey_adm"]
