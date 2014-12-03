
FROM softinnov/mininginx
MAINTAINER Ghislain Guiot

ADD nginx.conf /etc/nginx/
ADD client /client

VOLUME /var/log/nginx
VOLUME /etc/ssl/private

EXPOSE 8000

CMD ["nginx"]
