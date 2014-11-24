
FROM busybox:ubuntu-14.04
MAINTAINER Ghislain Guiot

ADD run.sh /run.sh

RUN chmod +x run.sh

ADD bearded-basket /exe/bearded-basket

EXPOSE 8002

CMD ["/run.sh"]
