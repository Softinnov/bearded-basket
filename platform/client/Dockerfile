
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot

RUN apt-get update -y
RUN apt-get install -y nginx

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/
ADD client /client

RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx"]
