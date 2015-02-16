
FROM jpetazzo/dind

ENV DEBIAN_FRONTEND noninteractive
#- custom from google/golang
RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git mercurial bzr
RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1
RUN mkdir -p /gopath/bin /gopath/pkg /gopath/src
RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
#-

WORKDIR /gopath/src/github.com/Softinnov/bearded-basket/

CMD ["bash"]
