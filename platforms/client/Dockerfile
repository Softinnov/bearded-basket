
FROM softinnov/nginx
MAINTAINER Ghislain Guiot

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/
ADD server* /etc/ssl/

VOLUME /var/log/nginx

EXPOSE 8000 443

CMD ["nginx"]
