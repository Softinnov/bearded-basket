
FROM debian:7.7
MAINTAINER Ghislain Guiot

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD run.sh /run.sh

RUN chmod +x run.sh

EXPOSE 8002

CMD ["/run.sh"]
