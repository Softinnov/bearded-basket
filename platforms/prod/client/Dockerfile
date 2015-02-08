
FROM softinnov/nginx
MAINTAINER Ghislain Guiot

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ADD nginx.conf /etc/nginx/
ADD client /client
ADD index.html /etc/nginx/html/index.html

VOLUME /var/log/nginx
VOLUME /etc/ssl/private

EXPOSE 80 443

CMD ["nginx"]
