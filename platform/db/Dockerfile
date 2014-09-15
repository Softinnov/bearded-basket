
FROM ubuntu:14.04
MAINTAINER Ghislain Guiot <gt.ghislain@gmail.com>

# Install MySQL.
RUN \
  apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server && \
  rm -rf /var/lib/apt/lists/* && \
  sed -i 's/^\(bind-address\s.*\)/# \1/' /etc/mysql/my.cnf && \
  sed -i 's/^\(log_error\s.*\)/# \1/' /etc/mysql/my.cnf && \
  echo "mysqld_safe &" > /tmp/config && \
  echo "mysqladmin --silent --wait=30 ping || exit 1" >> /tmp/config && \
  echo "mysql -e 'GRANT ALL PRIVILEGES ON *.* TO \"root\"@\"%\" WITH GRANT OPTION;'" >> /tmp/config && \
  bash /tmp/config && \
  rm -f /tmp/config

RUN mkdir data
ADD pdv.sql    data/pdv.sql
ADD pdv.txt.gz data/pdv.txt.gz

RUN /usr/sbin/mysqld & \
  sleep 10s && \
  mysql -uroot -e 'CREATE DATABASE prod' && \
  gunzip -f data/*.gz && \
  cat data/*.sql | mysql -u root prod && \
  mysqlimport -u root --local prod data/*.txt
  
# Expose ports.
EXPOSE 3306

# Define default command.
CMD ["mysqld_safe"]
