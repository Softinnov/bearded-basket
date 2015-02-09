FROM debian:7.7

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

RUN echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections
RUN dpkg --add-architecture i386
RUN apt-get update -qq && \
    apt-get -y install libc6:i386 libreadline6:i386 && \
    apt-get -y install curl runit && \
    rm -rf /var/lib/apt/lists/*

ENV CT_URL https://github.com/hashicorp/consul-template/releases/download/v0.5.1/consul-template_0.5.1_linux_amd64.tar.gz
RUN curl -L $CT_URL | tar -C /usr/local/bin --strip-components 1 -zxf -

ADD chey.service /etc/service/chey/run
ADD consul-template.service /etc/service/consul-template/run

ADD cheyenne-linux /cheyenne
ADD httpd.cfg /etc/consul-templates/httpd.cfg
ADD mysql-protocol.r /libs/mysql-protocol.r

VOLUME /esc-pdv
VOLUME /var/log

EXPOSE 80

CMD ["/usr/bin/runsvdir", "/etc/service"]
