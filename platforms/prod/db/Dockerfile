FROM debian:7.7
MAINTAINER Ghislain Guiot

# From tutum-docker-mysql
# Available at: https://github.com/tutumcloud/tutum-docker-mysql

# Install packages
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && \
  apt-get -yq install mysql-server pwgen && \
  rm -rf /var/lib/apt/lists/*

ADD update_dbdata.sh /

# Remove pre-installed database
RUN rm -rf /var/lib/mysql/*

# Add MySQL configuration
ADD my.master.cnf /etc/mysql/conf.d/my.cnf
ADD mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf

# Add VOLUMEs to allow backup of config and databases
VOLUME ["/var/lib/mysql", "/data"]

EXPOSE 3306

CMD ["mysqld_safe"]
