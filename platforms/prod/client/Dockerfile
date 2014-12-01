
FROM softinnov/mininginx
MAINTAINER Ghislain Guiot

ADD nginx.conf /etc/nginx/
ADD client /client

VOLUME /var/log/nginx

EXPOSE 8000

CMD ["nginx"]
