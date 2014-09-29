
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot

ADD bearded-basket bearded-basket
 
EXPOSE 8000
 
CMD /bearded-basket -db "$DBUSER:$DBPASS@(db:$DB_PORT_3306_TCP_PORT)/prod"
