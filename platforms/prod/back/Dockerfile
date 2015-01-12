
FROM debian:7.7
MAINTAINER Ghislain Guiot

RUN apt-get update && \
  apt-get install -y net-tools
RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD run.sh /run.sh

ADD bearded-basket /exe/bearded-basket

EXPOSE 8002

CMD ["/run.sh"]
