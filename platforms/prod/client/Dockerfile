
FROM softinnov/mininginx
MAINTAINER Ghislain Guiot

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/
ADD client /client

VOLUME /var/log/nginx

EXPOSE 8000

CMD ["nginx"]
