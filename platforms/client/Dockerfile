
FROM nginx:1.7
MAINTAINER Ghislain Guiot

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

RUN DEBIAN_FRONTEND=noninteractive \
    apt-get update -qq && \
    apt-get -y install curl runit && \
    rm -rf /var/lib/apt/lists/*

ENV CT_URL https://github.com/hashicorp/consul-template/releases/download/v0.5.1/consul-template_0.5.1_linux_amd64.tar.gz
RUN curl -L $CT_URL | tar -C /usr/local/bin --strip-components 1 -zxf -

ADD nginx.service /etc/service/nginx/run
ADD consul-template.service /etc/service/consul-template/run

RUN rm -v /etc/nginx/conf.d/*
ADD nginx.conf /etc/consul-templates/nginx.conf

ADD server* /etc/ssl/

VOLUME /var/log/nginx

EXPOSE 8000 443

CMD ["/usr/bin/runsvdir", "/etc/service"]
