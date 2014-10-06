
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y
RUN apt-get install -y nginx

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/
ADD run.sh /run.sh

RUN chmod +x run.sh

RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf

VOLUME /var/log/nginx

EXPOSE 8000

CMD ["/run.sh"]
