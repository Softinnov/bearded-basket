
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot <gt.ghislain@gmail.com>

RUN apt-get update
RUN apt-get install -y nginx

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/

RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx"]
