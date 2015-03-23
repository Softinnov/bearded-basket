FROM softinnov/cheyenne

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD httpd.pdv.cfg /httpd.cfg

RUN mv /cheyenne /chey_pdv

ADD esc-pdv/src /esc-pdv
VOLUME /var/log

EXPOSE 80

CMD ["/chey_pdv"]
