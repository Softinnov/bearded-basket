FROM debian:7.7

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections
RUN dpkg --add-architecture i386
RUN apt-get update -y
RUN apt-get -y install libc6:i386 libreadline6:i386

ADD cheyenne-linux /cheyenne
ADD httpd.cfg /httpd.cfg
ADD mysql-protocol.r /libs/mysql-protocol.r

VOLUME /esc-pdv
VOLUME /var/log

EXPOSE 80

CMD ["/cheyenne"]
