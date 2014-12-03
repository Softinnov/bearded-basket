FROM softinnov/prod-db
MAINTAINER Ghislain Guiot

ADD run.sh /

# Add MySQL configuration
ADD my.slave.cnf /etc/mysql/conf.d/my.cnf

CMD ["/run.sh"]
