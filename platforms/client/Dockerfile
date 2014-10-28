
FROM softinnov/nginx
MAINTAINER Ghislain Guiot

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/
ADD run.sh /run.sh

RUN chmod +x run.sh

VOLUME /var/log/nginx

EXPOSE 8000

CMD ["/run.sh"]
