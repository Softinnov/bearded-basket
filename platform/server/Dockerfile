
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot <gt.ghislain@gmail.com>

ADD bearded-basket bearded-basket
 
EXPOSE 8000
 
CMD /bearded-basket -db "root:@(db:$DB_PORT_3306_TCP_PORT)/prod"
