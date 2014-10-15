FROM ubuntu:14.04
MAINTAINER Ghislain Guiot

# From tutum-docker-mysql
# Available at: https://github.com/tutumcloud/tutum-docker-mysql

# Install packages
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && \
  apt-get -yq install mysql-server pwgen && \
  rm -rf /var/lib/apt/lists/*

# Remove pre-installed database
RUN rm -rf /var/lib/mysql/*

# Add MySQL configuration
ADD my.cnf /etc/mysql/conf.d/my.cnf
ADD mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf

# Add tables
ADD data/role.sql        /data/role.sql
ADD data/role.txt        /data/role.txt

ADD data/utilisateur.sql /data/utilisateur.sql
ADD data/utilisateur.txt /data/utilisateur.txt

ADD data/pdv.sql         /data/pdv.sql
ADD data/pdv.txt         /data/pdv.txt

# Add script
ADD run.sh /run.sh
RUN chmod 755 /*.sh

# Exposed ENV
ENV MYSQL_USER admin
ENV MYSQL_PASS admin

RUN /run.sh

# Add VOLUMEs to allow backup of config and databases
VOLUME ["/etc/mysql", "/var/lib/mysql"]

EXPOSE 3306

CMD ["mysqld_safe"]
